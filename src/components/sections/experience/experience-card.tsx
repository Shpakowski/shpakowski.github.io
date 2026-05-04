import { useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { motion, AnimatePresence } from 'framer-motion';
import { ChevronDown } from 'lucide-react';
import { useLayout } from '@/context';
import type { ExperienceEntry } from '@/types';
import { StackPills } from '@/components/ui';
import { AchievementItem } from '@/components/ui';
import { SubProjectCard } from './sub-project-card';

export function ExperienceCard({ entry, index }: { entry: ExperienceEntry; index: number }) {
  const { activeExperienceIndex, forceExpandSignal, openExperience } = useLayout();
  const expanded = activeExperienceIndex === index;
  const { t } = useTranslation();

  const hasSubProjects = entry.subProjects && entry.subProjects.length > 0;
  const hasDetails = entry.description || (entry.achievements && entry.achievements.length > 0);

  // 🎓 React Trend: Управляем стейтом (скроллом) через эффекты, подписанные на глобальный стейт Context API,
  // а не через грязное "прослушивание" window событий.
  useEffect(() => {
    if (expanded) {
      setTimeout(() => {
        const el = document.getElementById(`experience-${index}`);
        if (el) {
          el.scrollIntoView({ behavior: 'smooth', block: 'start' });
        }
      }, 320);
    }
  }, [expanded, index]);

  const handleToggle = () => {
    if (!expanded) {
      openExperience(index);
    } else {
      openExperience(null);
    }
  };

  return (
    <motion.div
      id={`experience-${index}`}
      initial={{ opacity: 0, y: 16 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ delay: 0.1 + index * 0.1, duration: 0.4 }}
      className="relative scroll-mt-24"
    >
      {/* Timeline dot */}
      <div className="absolute left-0 top-[15px] w-3 h-3 rounded-full bg-[var(--card)] border-2 border-[var(--timeline-dot)] z-10" />
      {/* Timeline line */}
      {index < 3 && (
        <div className="absolute left-[5px] top-[27px] bottom-0 w-0.5 bg-[var(--timeline-line)]" />
      )}

      <div className="ml-7">
        {/* Company header */}
        {/* 🎓 React Trend: Семантика и Доступность (A11y). Интерактивные элементы, 
            раскрывающие контент, должны быть <button>, а не <div>. Добавлен aria-expanded 
            для корректной работы со скринридерами (WCAG Guidelines). */}
        <button
          onClick={handleToggle}
          className="w-full text-left cursor-pointer flex items-start justify-between gap-3 py-2 group focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary rounded-md"
          aria-expanded={expanded}
        >
          <div className="min-w-0">
            <div className="flex items-baseline gap-2 flex-wrap">
              <h3 className="text-base font-bold text-[var(--foreground)]">
                {entry.company}
              </h3>
              <span className="text-xs font-mono text-[var(--muted)] whitespace-nowrap">
                {entry.dates}
              </span>
            </div>
            <p className="text-sm font-medium text-blue-600 dark:text-blue-400 mt-0.5">
              {entry.title}
            </p>
            {entry.location && (
              <p className="text-xs text-[var(--muted)] mt-0.5">{entry.location}</p>
            )}
          </div>

          {(hasSubProjects || hasDetails) && (
            <motion.div
              animate={{ rotate: expanded ? 180 : 0 }}
              transition={{ duration: 0.2 }}
              className="mt-2 shrink-0"
            >
              <ChevronDown className="h-6 w-6 text-[var(--timeline-dot)]" />
            </motion.div>
          )}
        </button>

        {/* Expanded content */}
        <AnimatePresence>
          {expanded && (
            <motion.div
              initial={{ height: 0, opacity: 0 }}
              animate={{ height: 'auto', opacity: 1 }}
              exit={{ height: 0, opacity: 0 }}
              transition={{ duration: 0.25 }}
              className="overflow-hidden"
            >
              {/* Role badge if present */}
              {entry.role && (
                <p className="text-xs text-[var(--muted)] mb-2">
                  {t('labels.projectRole')}: {entry.role}
                </p>
              )}

              {/* Description */}
              {entry.description && (
                <p className="text-sm text-[var(--muted)] leading-relaxed mb-3">
                  {entry.description}
                </p>
              )}

              {/* Achievements */}
              {entry.achievements && entry.achievements.length > 0 && (
                <ul className="space-y-2 mb-3">
                  {entry.achievements.map((a, i) => (
                    <AchievementItem key={i} text={a} />
                  ))}
                </ul>
              )}

              {/* Stack */}
              {entry.stack && <StackPills stack={entry.stack} />}

              {/* Sub-projects (EPAM) */}
              {hasSubProjects && (
                <div className="mt-4 space-y-1">
                  {entry.subProjects!.map((sp, i) => (
                    <SubProjectCard key={sp.title} project={sp} index={i} forceOpenSignal={expanded ? forceExpandSignal : 0} />
                  ))}
                </div>
              )}

              <div className="h-4" />
            </motion.div>
          )}
        </AnimatePresence>
      </div>
    </motion.div>
  );
}

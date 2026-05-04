import { useState, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { motion, AnimatePresence } from 'framer-motion';
import { ChevronDown } from 'lucide-react';

interface SubProject {
  title: string;
  role: string;
  description: string;
  achievements?: string[];
  stack: string;
}

interface ExperienceEntry {
  id: string;
  company: string;
  title: string;
  dates: string;
  location?: string;
  role?: string;
  description?: string;
  achievements?: string[];
  stack?: string;
  subProjects?: SubProject[];
}

import { SKILLS_CONFIG } from '../../config/skills.config';

function StackPills({ stack }: { stack: string }) {
  const getThemeForSkill = (skillName: string) => {
    const lowerSkill = skillName.toLowerCase();
    
    for (let i = 0; i < SKILLS_CONFIG.length; i++) {
      const categorySkills = SKILLS_CONFIG[i].items.map(s => s.toLowerCase());
      
      // Exact match
      if (categorySkills.includes(lowerSkill)) return SKILLS_CONFIG[i].theme;
      
      // Partial match
      for (const catSkill of categorySkills) {
        if (catSkill.length > 2 && (catSkill.includes(lowerSkill) || lowerSkill.includes(catSkill))) {
          return SKILLS_CONFIG[i].theme;
        }
      }
    }
    
    // Fallback dictionary for known terms in experience but not explicitly in skills
    const fallbackMap: Record<string, number> = {
      'go': 2, 'redis': 2, 'typeorm': 2, 'stripe': 2, 'rabbitmq': 2,
      'cloudflare': 3, 'gcp': 3, 'nx': 1, 'vite': 1, 'webpack': 1, 
      'telegram sdk': 1, 'rxjs': 1
    };
    
    if (fallbackMap[lowerSkill] !== undefined) {
      return SKILLS_CONFIG[fallbackMap[lowerSkill]].theme;
    }

    return null; // Fallback to default gray
  };

  const items = stack.split(',').map(s => s.trim().replace(/\.$/, '')).filter(Boolean);
  
  return (
    <div className="flex flex-wrap gap-1.5 mt-3">
      {items.map((item) => {
        const theme = getThemeForSkill(item);
        
        if (theme) {
          return (
            <motion.span
              key={item}
              whileHover={{ scale: 1.05 }}
              className={`
                inline-flex items-center px-2.5 py-0.5 rounded-md text-[11px] font-medium
                border transition-all duration-200 cursor-default
                ${theme.pill} ${theme.text} ${theme.glow}
              `}
            >
              {item}
            </motion.span>
          );
        }
        
        return (
          <span
            key={item}
            className="
              inline-flex items-center px-2.5 py-0.5 rounded-md text-[11px] font-medium
              bg-[var(--pill-bg)] text-[var(--muted)] border border-[var(--pill-border)]
              transition-all duration-200 hover:shadow-sm cursor-default
            "
          >
            {item}
          </span>
        );
      })}
    </div>
  );
}

function AchievementItem({ text }: { text: string }) {
  // Parse **bold** markdown
  const parts = text.split(/\*\*(.*?)\*\*/g);
  return (
    <li className="text-sm text-[var(--muted)] leading-relaxed pl-4 relative before:content-[''] before:absolute before:left-0 before:top-[9px] before:w-1.5 before:h-1.5 before:rounded-full before:bg-[var(--timeline-dot)]/40">
      {parts.map((part, i) =>
        i % 2 === 1 ? (
          <strong key={i} className="font-semibold text-[var(--foreground)]">{part}</strong>
        ) : (
          <span key={i}>{part}</span>
        )
      )}
    </li>
  );
}

function SubProjectCard({ project, index, forceOpenSignal }: { project: SubProject; index: number; forceOpenSignal: number }) {
  const [open, setOpen] = useState(forceOpenSignal > 0);

  const [prevSignal, setPrevSignal] = useState(forceOpenSignal);

  if (forceOpenSignal !== prevSignal) {
    setPrevSignal(forceOpenSignal);
    if (forceOpenSignal > 0) {
      setOpen(true);
    }
  }

  return (
    <motion.div
      initial={{ opacity: 0, x: -8 }}
      animate={{ opacity: 1, x: 0 }}
      transition={{ delay: index * 0.08, duration: 0.3 }}
      className="
        relative ml-5 pl-5
        border-l-2 border-[var(--border)]
      "
    >
      <div
        onClick={() => setOpen(!open)}
        className="
          cursor-pointer py-3 group
          flex items-start justify-between gap-3
        "
      >
        <div className="min-w-0">
          <h4 className="text-sm font-semibold text-[var(--foreground)] group-hover:text-blue-600 dark:group-hover:text-blue-400 transition-colors">
            {project.title}
          </h4>
          <p className="text-xs text-[var(--muted)] mt-0.5">{project.role}</p>
        </div>
        <motion.div
          animate={{ rotate: open ? 180 : 0 }}
          transition={{ duration: 0.2 }}
          className="mt-1 shrink-0"
        >
          <ChevronDown className="h-6 w-6 text-[var(--timeline-dot)]" />
        </motion.div>
      </div>

      <AnimatePresence>
        {open && (
          <motion.div
            initial={{ height: 0, opacity: 0 }}
            animate={{ height: 'auto', opacity: 1 }}
            exit={{ height: 0, opacity: 0 }}
            transition={{ duration: 0.25 }}
            className="overflow-hidden"
          >
            <p className="text-sm text-[var(--muted)] leading-relaxed mb-3">
              {project.description}
            </p>
            {project.achievements && (
              <ul className="space-y-2 mb-3">
                {project.achievements.map((a, i) => (
                  <AchievementItem key={i} text={a} />
                ))}
              </ul>
            )}
            <StackPills stack={project.stack} />
          </motion.div>
        )}
      </AnimatePresence>
    </motion.div>
  );
}

function ExperienceCard({ entry, index }: { entry: ExperienceEntry; index: number }) {
  const [expanded, setExpanded] = useState(false); // All closed by default
  const [forceOpenSignal, setForceOpenSignal] = useState(0);
  const { t } = useTranslation();

  const hasSubProjects = entry.subProjects && entry.subProjects.length > 0;
  const hasDetails = entry.description || (entry.achievements && entry.achievements.length > 0);

  useEffect(() => {
    const handleOpenExperience = (e: Event) => {
      const customEvent = e as CustomEvent<{ index: number, expandSubProjects?: boolean }>;
      if (customEvent.detail.index === index) {
        setExpanded(true);
        if (customEvent.detail.expandSubProjects) {
          setForceOpenSignal(prev => prev + 1);
        } else {
          setForceOpenSignal(0);
        }
        // Wait for accordion animations to settle before scrolling
        setTimeout(() => {
          const el = document.getElementById(`experience-${index}`);
          if (el) {
            el.scrollIntoView({ behavior: 'smooth', block: 'start' });
          }
        }, 320);
      } else {
        setExpanded(false);
        setForceOpenSignal(0);
      }
    };
    window.addEventListener('open-experience', handleOpenExperience);
    return () => window.removeEventListener('open-experience', handleOpenExperience);
  }, [index]);

  const handleToggle = () => {
    if (!expanded) {
      const event = new CustomEvent('open-experience', { detail: { index } });
      window.dispatchEvent(event);
    } else {
      setExpanded(false);
      const event = new CustomEvent('open-experience', { detail: { index: -1 } });
      window.dispatchEvent(event);
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
        <div
          onClick={handleToggle}
          className="cursor-pointer flex items-start justify-between gap-3 py-2 group"
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
        </div>

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
                    <SubProjectCard key={sp.title} project={sp} index={i} forceOpenSignal={forceOpenSignal} />
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

export function ExperienceSection() {
  const { t } = useTranslation();
  const experience = t('experience', { returnObjects: true }) as ExperienceEntry[];

  return (
    <section id="experience-section" className="py-2">
      {/* Section label */}
      <div className="flex items-center gap-3 mb-6">
        <h2 className="text-xs font-semibold uppercase tracking-[0.15em] text-[var(--muted)]">
          {t('sections.experience')}
        </h2>
        <div className="flex-1 h-px bg-[var(--border)]" />
      </div>

      {/* Timeline */}
      <div className="relative space-y-6 pl-1">
        {experience.map((entry, i) => (
          <ExperienceCard key={entry.id} entry={entry} index={i} />
        ))}
      </div>
    </section>
  );
}

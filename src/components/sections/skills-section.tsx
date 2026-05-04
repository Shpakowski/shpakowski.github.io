import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { motion, AnimatePresence } from 'framer-motion';
import { SKILLS_CONFIG } from '@/config';
import { SectionHeader } from '@/components/ui';

export function SkillsSection() {
  const { t } = useTranslation();
  const [activeTab, setActiveTab] = useState(0);

  const activeCategory = SKILLS_CONFIG[activeTab] ?? SKILLS_CONFIG[0];
  const theme = activeCategory.theme;
  const skillItems = activeCategory.items;

  return (
    <section id="skills-section" className="py-2">
      <SectionHeader title={t('sections.skills')} />

      {/* Tab bar — horizontal scrollable */}
      {/* 🎓 React Trend: A11y (Accessibility). Для компонентов типа "вкладки" обязательно указываем 
          role="tablist" и role="tab", чтобы скринридеры понимали навигационную структуру интерфейса. */}
      <div className="relative mb-6">
        <div role="tablist" className="flex gap-1.5 overflow-x-auto scrollbar-none pb-1 -mx-1 px-1">
          {SKILLS_CONFIG.map((category, index) => {
            const isActive = index === activeTab;
            const Icon = category.icon;
            const label = category.name;

            return (
              <button
                key={category.id}
                role="tab"
                aria-selected={isActive}
                aria-controls={`panel-${category.id}`}
                onClick={() => setActiveTab(index)}
                className={`
                  relative flex items-center gap-2 px-4 py-2.5 rounded-xl text-sm font-medium
                  whitespace-nowrap transition-all duration-200
                  ${isActive
                    ? `${category.theme.bg} ${category.theme.text} shadow-sm`
                    : 'text-[var(--muted)] hover:text-[var(--foreground)] hover:bg-[var(--card)]'
                  }
                `}
              >
                {isActive && (
                  <motion.div
                    layoutId="skill-tab-indicator"
                    className={`absolute inset-0 rounded-xl ${category.theme.bg} border ${category.theme.border}/30`}
                    transition={{ type: 'spring', stiffness: 400, damping: 30 }}
                  />
                )}
                <Icon className="relative z-10 h-4 w-4" />
                <span className="relative z-10">{label}</span>
              </button>
            );
          })}
        </div>
      </div>

      {/* Skill pills */}
      <AnimatePresence mode="wait">
        <motion.div
          key={activeTab}
          id={`panel-${activeCategory.id}`}
          role="tabpanel"
          initial={{ opacity: 0, y: 8 }}
          animate={{ opacity: 1, y: 0 }}
          exit={{ opacity: 0, y: -8 }}
          transition={{ duration: 0.2 }}
          className="flex flex-wrap gap-2"
        >
          {skillItems.map((skill, i) => (
            <motion.span
              key={skill}
              initial={{ opacity: 0, scale: 0.92 }}
              animate={{ opacity: 1, scale: 1 }}
              transition={{ delay: i * 0.03, duration: 0.2 }}
              className={`
                inline-flex items-center px-3.5 py-1.5 rounded-lg text-sm font-medium
                border transition-all duration-200
                ${theme.pill} ${theme.text} ${theme.glow}
              `}
            >
              {skill}
            </motion.span>
          ))}
        </motion.div>
      </AnimatePresence>
    </section>
  );
}

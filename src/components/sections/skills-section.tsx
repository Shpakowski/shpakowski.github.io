import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { motion, AnimatePresence } from 'framer-motion';
import { BrainCircuit, Crown, Monitor, Server, Cloud } from 'lucide-react';
import type { LucideIcon } from 'lucide-react';

interface SkillCategory {
  category: string;
  items: string;
}

/* ---------- Color & icon config by INDEX (language-independent) ---------- */

export interface TabTheme {
  bg: string;
  text: string;
  pill: string;
  glow: string;
  border: string;
  icon: LucideIcon;
  shortName: Record<string, string>;   // keyed by i18n language code
}

export const TAB_THEMES: TabTheme[] = [
  // 0 — AI-Driven Development
  {
    bg: 'bg-rose-500/10 dark:bg-rose-500/15',
    text: 'text-rose-700 dark:text-rose-300',
    pill: 'bg-rose-100 dark:bg-rose-900/40 border-rose-200 dark:border-rose-700/50',
    glow: 'hover:shadow-[0_0_12px_rgba(244,63,94,0.15)]',
    border: 'border-rose-500',
    icon: BrainCircuit,
    shortName: { en: 'AI-Driven', ru: 'AI-Driven' },
  },
  // 1 — Frontend
  {
    bg: 'bg-amber-500/10 dark:bg-amber-500/15',
    text: 'text-amber-700 dark:text-amber-300',
    pill: 'bg-amber-100 dark:bg-amber-900/40 border-amber-200 dark:border-amber-700/50',
    glow: 'hover:shadow-[0_0_12px_rgba(245,158,11,0.15)]',
    border: 'border-amber-500',
    icon: Monitor,
    shortName: { en: 'Frontend', ru: 'Frontend' },
  },
  // 2 — Backend
  {
    bg: 'bg-emerald-500/10 dark:bg-emerald-500/15',
    text: 'text-emerald-700 dark:text-emerald-300',
    pill: 'bg-emerald-100 dark:bg-emerald-900/40 border-emerald-200 dark:border-emerald-700/50',
    glow: 'hover:shadow-[0_0_12px_rgba(16,185,129,0.15)]',
    border: 'border-emerald-500',
    icon: Server,
    shortName: { en: 'Backend', ru: 'Backend' },
  },
  // 3 — DevOps
  {
    bg: 'bg-blue-500/10 dark:bg-blue-500/15',
    text: 'text-blue-700 dark:text-blue-300',
    pill: 'bg-blue-100 dark:bg-blue-900/40 border-blue-200 dark:border-blue-700/50',
    glow: 'hover:shadow-[0_0_12px_rgba(59,130,246,0.15)]',
    border: 'border-blue-500',
    icon: Cloud,
    shortName: { en: 'DevOps', ru: 'DevOps' },
  },
  // 4 — Leadership
  {
    bg: 'bg-violet-500/10 dark:bg-violet-500/15',
    text: 'text-violet-700 dark:text-violet-300',
    pill: 'bg-violet-100 dark:bg-violet-900/40 border-violet-200 dark:border-violet-700/50',
    glow: 'hover:shadow-[0_0_12px_rgba(139,92,246,0.15)]',
    border: 'border-violet-500',
    icon: Crown,
    shortName: { en: 'Tech Lead', ru: 'Tech Lead' },
  },
];

export function SkillsSection() {
  const { t, i18n } = useTranslation();
  const skills = t('skills', { returnObjects: true }) as SkillCategory[];
  const [activeTab, setActiveTab] = useState(0);
  const lang = i18n.language?.startsWith('ru') ? 'ru' : 'en';

  const theme = TAB_THEMES[activeTab] ?? TAB_THEMES[0];
  const skillItems = skills[activeTab].items.split(',').map(s => s.trim()).filter(Boolean);

  return (
    <section id="skills-section" className="py-2">
      {/* Section label */}
      <div className="flex items-center gap-3 mb-6">
        <h2 className="text-xs font-semibold uppercase tracking-[0.15em] text-[var(--muted)]">
          {t('sections.skills')}
        </h2>
        <div className="flex-1 h-px bg-[var(--border)]" />
      </div>

      {/* Tab bar — horizontal scrollable */}
      <div className="relative mb-6">
        <div className="flex gap-1.5 overflow-x-auto scrollbar-none pb-1 -mx-1 px-1">
          {skills.map((_, index) => {
            const tabTheme = TAB_THEMES[index] ?? TAB_THEMES[0];
            const isActive = index === activeTab;
            const Icon = tabTheme.icon;
            const label = tabTheme.shortName[lang] ?? tabTheme.shortName.en;

            return (
              <button
                key={index}
                onClick={() => setActiveTab(index)}
                className={`
                  relative flex items-center gap-2 px-4 py-2.5 rounded-xl text-sm font-medium
                  whitespace-nowrap transition-all duration-200
                  ${isActive
                    ? `${tabTheme.bg} ${tabTheme.text} shadow-sm`
                    : 'text-[var(--muted)] hover:text-[var(--foreground)] hover:bg-[var(--card)]'
                  }
                `}
              >
                {isActive && (
                  <motion.div
                    layoutId="skill-tab-indicator"
                    className={`absolute inset-0 rounded-xl ${tabTheme.bg} border ${tabTheme.border}/30`}
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

import { useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { ChevronDown } from 'lucide-react';
import type { SubProject } from '@/types';
import { StackPills } from '@/components/ui';
import { AchievementItem } from '@/components/ui';

export function SubProjectCard({ project, index, forceOpenSignal }: { project: SubProject; index: number; forceOpenSignal: number }) {
  const [open, setOpen] = useState(forceOpenSignal > 0);
  const [prevSignal, setPrevSignal] = useState(forceOpenSignal);

  // 🎓 React Trend: Обновление состояния при изменении пропсов (Derived State) 
  // по современным стандартам (React 18+) делается прямо в теле рендера, а не через useEffect.
  // Это предотвращает каскадные ререндеры (eslint: react-hooks/set-state-in-effect).
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
      {/* 🎓 React Trend: Интерактивные элементы должны быть button для A11y (Доступности). */}
      <button
        onClick={() => setOpen(!open)}
        className="
          cursor-pointer py-3 group w-full text-left
          flex items-start justify-between gap-3
          focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary rounded-md
        "
        aria-expanded={open}
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
      </button>

      <AnimatePresence>
        {open && (
          <motion.div
            initial={{ height: 0, opacity: 0 }}
            animate={{ height: 'auto', opacity: 1 }}
            exit={{ height: 0, opacity: 0 }}
            transition={{ duration: 0.25 }}
            className="overflow-hidden"
          >
            <p className="text-sm text-[var(--muted)] leading-relaxed mb-3 mt-2">
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

import * as React from 'react';
import { cn } from '@/utils';

interface SectionHeaderProps extends React.HTMLAttributes<HTMLDivElement> {
  title: string;
}

// 🎓 React Trend: Создание переиспользуемых базовых UI-компонентов (Атомарный дизайн).
// Это убирает дублирование HTML/Tailwind классов по всему проекту (DRY).
export function SectionHeader({ title, className, ...props }: SectionHeaderProps) {
  return (
    <div className={cn("flex items-center gap-3 mb-6", className)} {...props}>
      <h2 className="text-xs font-semibold uppercase tracking-[0.15em] text-[var(--muted)]">
        {title}
      </h2>
      <div className="flex-1 h-px bg-[var(--border)]" />
    </div>
  );
}

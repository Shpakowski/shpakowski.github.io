import { motion } from 'framer-motion';
import { getThemeForSkill } from '@/utils';

export function StackPills({ stack }: { stack: string }) {
  const items = stack.split(',').map(s => s.trim().replace(/\.$/, '').replace(/\s+/g, ' ')).filter(Boolean);
  
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

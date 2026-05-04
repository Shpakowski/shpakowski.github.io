import { motion, useInView } from 'framer-motion';
import { useRef } from 'react';
import { useCountUp } from '@/hooks/useCountUp';

// 🎓 React Trend: Логика таймеров вынесена в хук useCountUp. Компонент стал "тонким" и отвечает только за UI.
function CountUp({ target, suffix, delay }: { target: number; suffix: string; delay: number }) {
  const count = useCountUp(target, delay);
  return <span>{count}{suffix}</span>;
}

export function AnimatedMetric({ value, delay }: { value: string; delay: number }) {
  const ref = useRef<HTMLDivElement>(null);
  
  // 🎓 React Trend: Используем готовые хуки (useInView из framer-motion) вместо ручного IntersectionObserver.
  // Это избавляет от бойлерплейт-кода, автоматически чистит подписки и снижает риск утечек памяти (DRY & KISS).
  const visible = useInView(ref, { once: true, amount: 0.3 });

  // If the value is a pure number like "15+", animate count-up
  const numMatch = value.match(/^(\d+)(.*)$/);

  if (numMatch && visible) {
    return (
      <div ref={ref}>
        <CountUp target={parseInt(numMatch[1])} suffix={numMatch[2]} delay={delay} />
      </div>
    );
  }

  return (
    <div ref={ref} className={visible ? 'opacity-100' : 'opacity-0'}>
      <motion.span
        initial={{ opacity: 0, y: 10 }}
        animate={visible ? { opacity: 1, y: 0 } : {}}
        transition={{ delay, duration: 0.4 }}
      >
        {value}
      </motion.span>
    </div>
  );
}

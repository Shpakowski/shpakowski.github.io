import { useTranslation } from 'react-i18next';
import { motion } from 'framer-motion';
import { useEffect, useRef, useState } from 'react';

interface Highlight {
  metric: string;
  title: string;
  description: string;
}

function AnimatedMetric({ value, delay }: { value: string; delay: number }) {
  const [visible, setVisible] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) setVisible(true);
      },
      { threshold: 0.3 }
    );
    if (ref.current) observer.observe(ref.current);
    return () => observer.disconnect();
  }, []);

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

function CountUp({ target, suffix, delay }: { target: number; suffix: string; delay: number }) {
  const [count, setCount] = useState(0);
  const [started, setStarted] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => setStarted(true), delay * 1000);
    return () => clearTimeout(timer);
  }, [delay]);

  useEffect(() => {
    if (!started) return;
    const duration = 1200;
    const steps = 30;
    const increment = target / steps;
    let current = 0;
    const interval = setInterval(() => {
      current += increment;
      if (current >= target) {
        setCount(target);
        clearInterval(interval);
      } else {
        setCount(Math.floor(current));
      }
    }, duration / steps);
    return () => clearInterval(interval);
  }, [started, target]);

  return <span>{count}{suffix}</span>;
}

export function HeroSection() {
  const { t } = useTranslation();
  const highlights = t('profile.highlights', { returnObjects: true }) as Highlight[];

  return (
    <section id="hero-section" className="py-2">
      {/* Metric cards grid */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-3 mb-6">
        {highlights.map((h, i) => (
          <motion.div
            key={h.title}
            initial={{ opacity: 0, y: 16 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.15 + i * 0.1, duration: 0.4 }}
            className="
              group relative p-5 rounded-xl
              bg-[var(--card)] border border-[var(--border)]
              hover:border-blue-400/40 dark:hover:border-blue-400/30
              hover:shadow-[0_2px_20px_rgba(59,130,246,0.08)]
              transition-all duration-300
            "
          >
            <div className="text-3xl font-bold tracking-tight text-[var(--foreground)] mb-1">
              <AnimatedMetric value={h.metric} delay={0.3 + i * 0.15} />
            </div>
            <div className="text-[12px] font-bold uppercase tracking-wider text-[var(--foreground)] opacity-90 mb-2 whitespace-nowrap">
              {h.title}
            </div>
            <div className="text-[13px] text-[var(--muted)] leading-relaxed group-hover:text-[var(--foreground)] transition-colors duration-300">
              {h.description}
            </div>
          </motion.div>
        ))}
      </div>

      {/* Compact profile summary — below cards */}
      <motion.p
        initial={{ opacity: 0, y: 12 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5, delay: 0.5 }}
        className="text-sm leading-relaxed text-[var(--muted)] max-w-2xl mb-6"
        dangerouslySetInnerHTML={{ __html: t('profile.summary') }}
      />

      {/* CTA Banner */}
      <motion.div
        initial={{ opacity: 0, y: 8 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.7, duration: 0.4 }}
        className="
          flex items-center gap-3 p-4 rounded-xl
          bg-[var(--cta-bg)] border border-[var(--cta-border)]
        "
      >
        <span className="relative flex h-2.5 w-2.5 shrink-0">
          <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75" />
          <span className="relative inline-flex rounded-full h-2.5 w-2.5 bg-emerald-500" />
        </span>
        <p className="text-sm font-medium text-[var(--cta-text)]">
          {t('profile.target')}
        </p>
      </motion.div>
    </section>
  );
}

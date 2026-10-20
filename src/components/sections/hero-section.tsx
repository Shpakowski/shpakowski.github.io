import { useTranslation } from 'react-i18next';
import { motion } from 'framer-motion';
import type { Highlight } from '@/types';
import { AnimatedMetric } from '@/components/ui';

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
              bg-card border border-border
              hover:border-blue-400/40 dark:hover:border-blue-400/30
              hover:shadow-[0_2px_20px_rgba(59,130,246,0.08)]
              transition-all duration-300
            "
          >
            <div className="text-3xl font-bold tracking-tight text-foreground mb-1">
              <AnimatedMetric value={h.metric} delay={0.3 + i * 0.15} />
            </div>
            <div className="text-[12px] font-bold uppercase tracking-wider text-foreground opacity-90 mb-2 whitespace-nowrap">
              {h.title}
            </div>
            <div className="text-[13px] text-muted leading-relaxed group-hover:text-foreground transition-colors duration-300">
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
        className="text-sm leading-relaxed text-muted w-full whitespace-pre-line"
        dangerouslySetInnerHTML={{ __html: t('profile.summary') }}
      />


    </section>
  );
}

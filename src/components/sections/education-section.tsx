import { useTranslation } from 'react-i18next';
import { motion } from 'framer-motion';
import { SectionHeader } from '@/components/ui';

export interface EducationEntry {
  degree: string;
  dates?: string;
  details: string;
}

export function EducationSection() {
  const { t } = useTranslation();
  const education = t('education', { returnObjects: true }) as EducationEntry[];

  if (!Array.isArray(education) || education.length === 0) {
    return null;
  }

  return (
    <section id="education-section" className="py-2">
      <SectionHeader title={t('sections.education')} />

      <div className="space-y-4 pl-1">
        {education.map((item, i) => (
          <motion.div
            key={item.degree}
            initial={{ opacity: 0, y: 16 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.1 + i * 0.1, duration: 0.4 }}
            className="relative"
          >
            <div className="absolute left-0 top-[15px] w-3 h-3 rounded-full bg-card border-2 border-timeline-dot z-10" />
            <div className="ml-7 py-2">
              <div className="flex items-baseline gap-2 flex-wrap">
                <h3 className="text-base font-bold text-foreground">
                  {item.degree}
                </h3>
                {item.dates && (
                  <span className="text-xs font-mono text-muted whitespace-nowrap">
                    {item.dates}
                  </span>
                )}
              </div>
              <p className="text-sm text-muted mt-0.5">
                {item.details}
              </p>
            </div>
          </motion.div>
        ))}
      </div>
    </section>
  );
}

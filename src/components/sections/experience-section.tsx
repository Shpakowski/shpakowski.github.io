import { useTranslation } from 'react-i18next';
import { SectionHeader } from '@/components/ui';
import type { ExperienceEntry } from '@/types';
import { ExperienceCard } from './experience/experience-card';

export function ExperienceSection() {
  const { t } = useTranslation();
  const experience = t('experience', { returnObjects: true }) as ExperienceEntry[];

  return (
    <section id="experience-section" className="py-2">
      <SectionHeader title={t('sections.experience')} />

      {/* Timeline */}
      <div className="relative space-y-6 pl-1">
        {experience.map((entry, i) => (
          <ExperienceCard key={entry.id} entry={entry} index={i} />
        ))}
      </div>
    </section>
  );
}

import { useTranslation } from 'react-i18next';
import { Card, CardContent } from '../ui/card';
import { SectionBlock } from '../ui/section-block';

export function SkillsSection() {
  const { t } = useTranslation();
  
  // We expect an array of skills from translation.json
  const skills = t('skills', { returnObjects: true }) as Array<{ category: string, items: string }>;

  if (!Array.isArray(skills)) return null;

  return (
    <SectionBlock id="skills" title={t('sections.skills')}>
      <Card>
        <CardContent className="p-6 space-y-4">
          {skills.map((skillGroup, index) => (
            <div key={index} className="space-y-1 text-sm md:text-base">
              <span className="font-semibold text-foreground mr-2">
                {skillGroup.category}:
              </span>
              <span className="text-muted-foreground leading-relaxed">
                {skillGroup.items}
              </span>
            </div>
          ))}
        </CardContent>
      </Card>
    </SectionBlock>
  );
}

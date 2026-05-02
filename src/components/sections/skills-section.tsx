
import { useTranslation } from 'react-i18next';
import { Card, CardContent } from '../ui/card';

export function SkillsSection() {
  const { t } = useTranslation();
  
  // We expect an array of skills from translation.json
  const skills = t('skills', { returnObjects: true }) as Array<{ category: string, items: string }>;

  if (!Array.isArray(skills)) return null;

  return (
    <section id="skills" className="mb-12 scroll-mt-12">
      <h2 className="mb-6 text-2xl font-bold tracking-tight text-foreground">
        {t('sections.skills')}
      </h2>
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
    </section>
  );
}

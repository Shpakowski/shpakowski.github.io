
import { useTranslation } from 'react-i18next';
import { Card, CardContent } from '../ui/card';

export function ProfileSection() {
  const { t } = useTranslation();

  return (
    <section id="profile" className="mb-12 scroll-mt-12">
      <h2 className="mb-6 text-2xl font-bold tracking-tight text-foreground">
        {t('sections.profile')}
      </h2>
      <Card>
        <CardContent className="p-6">
          <div className="prose prose-sm md:prose-base dark:prose-invert max-w-none">
            {t('profile.description').split('\n').map((paragraph, index) => (
              paragraph.trim() ? <p key={index}>{paragraph}</p> : null
            ))}
          </div>
          
          <div className="mt-6 rounded-lg bg-primary/10 p-4 border border-primary/20">
            <p className="text-sm font-medium text-primary">
              {t('profile.target')}
            </p>
          </div>
        </CardContent>
      </Card>
    </section>
  );
}

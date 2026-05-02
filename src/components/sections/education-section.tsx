
import { useTranslation } from 'react-i18next';
import { Card, CardContent } from '../ui/card';

export function EducationSection() {
  const { t } = useTranslation();
  
  const certs = t('certifications', { returnObjects: true }) as Array<any>;
  const edu = t('education', { returnObjects: true }) as Array<any>;

  const hasCerts = Array.isArray(certs) && certs.length > 0;
  const hasEdu = Array.isArray(edu) && edu.length > 0;

  if (!hasCerts && !hasEdu) return null;

  return (
    <div className="space-y-12">
      {hasCerts && (
        <section id="certifications" className="scroll-mt-12">
          <h2 className="mb-6 text-2xl font-bold tracking-tight text-foreground">
            {t('sections.certifications')}
          </h2>
          <Card>
            <CardContent className="p-6 space-y-6">
              {certs.map((item, index) => (
                <div key={index}>
                  <h3 className="text-base font-bold text-foreground mb-1">{item.title}</h3>
                  <p className="text-sm text-muted-foreground">{item.details}</p>
                </div>
              ))}
            </CardContent>
          </Card>
        </section>
      )}

      {hasEdu && (
        <section id="education" className="scroll-mt-12">
          <h2 className="mb-6 text-2xl font-bold tracking-tight text-foreground">
            {t('sections.education')}
          </h2>
          <Card>
            <CardContent className="p-6 space-y-6">
              {edu.map((item, index) => (
                <div key={index}>
                  <h3 className="text-base font-bold text-foreground mb-1">{item.degree}</h3>
                  <p className="text-sm text-muted-foreground">{item.details}</p>
                </div>
              ))}
            </CardContent>
          </Card>
        </section>
      )}
    </div>
  );
}

import { useTranslation } from 'react-i18next';
import { Card, CardContent } from '../ui/card';
import { SectionBlock } from '../ui/section-block';

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
        <SectionBlock id="certifications" title={t('sections.certifications')}>
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
        </SectionBlock>
      )}

      {hasEdu && (
        <SectionBlock id="education" title={t('sections.education')}>
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
        </SectionBlock>
      )}
    </div>
  );
}

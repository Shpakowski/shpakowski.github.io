import { useTranslation } from 'react-i18next';
import { Card, CardContent } from '../ui/card';
import ReactMarkdown from 'react-markdown';
import { SectionBlock } from '../ui/section-block';

export function ExperienceSection() {
  const { t } = useTranslation();
  
  const experiences = t('experience', { returnObjects: true }) as Array<any>;

  if (!Array.isArray(experiences)) return null;

  return (
    <SectionBlock id="experience" title={t('sections.experience')}>
      <div className="space-y-6">
        {experiences.map((job) => (
          <Card key={job.id} className="overflow-hidden">
            <CardContent className="p-6 space-y-4">
              <div className="flex flex-col md:flex-row md:items-baseline md:justify-between gap-1">
                <h3 className="text-lg font-bold text-foreground">
                  {job.title}
                </h3>
                <span className="text-sm font-medium text-muted-foreground whitespace-nowrap">
                  {job.dates}
                </span>
              </div>
              
              <div>
                <div className="text-base font-semibold text-primary mb-1">
                  {job.company}
                </div>
                {job.location && (
                  <div className="text-xs text-muted-foreground mb-2">
                    {job.location}
                  </div>
                )}
                {job.role && (
                  <div className="text-sm font-medium text-foreground mb-3 border-l-2 border-primary/40 pl-3 py-0.5 bg-primary/5">
                    {t('labels.projectRole')}: {job.role}
                  </div>
                )}
              </div>

              {job.description && (
                <div className="prose prose-sm dark:prose-invert max-w-none text-muted-foreground">
                  {job.description.split('\n').map((paragraph: string, idx: number) => (
                    paragraph.trim() ? <p key={idx}>{paragraph}</p> : null
                  ))}
                </div>
              )}

              {job.achievements && (
                <ul className="list-disc pl-5 space-y-1.5 text-sm text-muted-foreground mt-3">
                  {job.achievements.map((item: string, idx: number) => (
                    <li key={idx}>
                      <ReactMarkdown 
                        components={{
                          strong: ({node, ...props}) => <strong className="font-semibold text-foreground" {...props} />
                        }}
                      >
                        {item}
                      </ReactMarkdown>
                    </li>
                  ))}
                </ul>
              )}

              {job.stack && (
                <div className="mt-4 pt-3 border-t border-border/50">
                  <span className="text-xs font-semibold text-foreground mr-2 uppercase tracking-wide">
                    {t('labels.stack')}:
                  </span>
                  <span className="text-sm text-muted-foreground">
                    {job.stack}
                  </span>
                </div>
              )}

              {/* Sub Projects (if any) */}
              {job.subProjects && job.subProjects.length > 0 && (
                <div className="mt-6 space-y-6 pt-4 border-t border-border/50">
                  {job.subProjects.map((sub: any, subIdx: number) => (
                    <div key={subIdx} className="pl-4 border-l border-border/70 relative">
                      <div className="absolute -left-1.5 top-2 h-3 w-3 rounded-full bg-border" />
                      
                      <h4 className="text-md font-bold text-foreground mb-1">
                        {sub.title}
                      </h4>
                      <div className="text-sm font-medium text-foreground mb-2">
                        {t('labels.projectRole')}: {sub.role}
                      </div>
                      
                      <p className="text-sm text-muted-foreground mb-3">
                        {sub.description}
                      </p>
                      
                      {sub.achievements && (
                        <ul className="list-disc pl-5 space-y-1.5 text-sm text-muted-foreground mb-3">
                          {sub.achievements.map((item: string, idx: number) => (
                            <li key={idx}>
                              <ReactMarkdown 
                                components={{
                                  strong: ({node, ...props}) => <strong className="font-semibold text-foreground" {...props} />
                                }}
                              >
                                {item}
                              </ReactMarkdown>
                            </li>
                          ))}
                        </ul>
                      )}
                      
                      <div className="text-sm text-muted-foreground">
                        <strong className="text-foreground">{t('labels.stack')}:</strong> {sub.stack}
                      </div>
                    </div>
                  ))}
                </div>
              )}
            </CardContent>
          </Card>
        ))}
      </div>
    </SectionBlock>
  );
}

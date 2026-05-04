import { useState, useEffect } from 'react';
import { profileConfig } from '../../config/profile.config';
import { Avatar } from '../ui/avatar';
import avatarImage from '../../assets/avatar.JPG';
import { useTranslation } from 'react-i18next';

export function Sidebar() {
  const { t } = useTranslation();
  const experience = t('experience', { returnObjects: true }) as Array<Record<string, string>>;
  const [activeIndex, setActiveIndex] = useState<number | null>(null);

  useEffect(() => {
    const handleOpenExperience = (e: Event) => {
      const customEvent = e as CustomEvent<{ index: number }>;
      setActiveIndex(customEvent.detail.index);
    };
    window.addEventListener('open-experience', handleOpenExperience);
    return () => window.removeEventListener('open-experience', handleOpenExperience);
  }, []);

  return (
    <aside className="flex flex-col gap-4 p-6 pt-0 md:p-8 md:pt-0 lg:sticky lg:top-0 lg:h-screen lg:w-80 xl:w-96 lg:overflow-y-auto bg-card/50">


      <div className="flex flex-col items-center text-center gap-2">
        <Avatar
          src={avatarImage}
          alt={profileConfig.fullName}
          className="h-[180px] w-[180px] shadow-md ring-2 ring-primary/20"
        />
        <div>
          <h1 className="text-2xl font-bold tracking-tight">{profileConfig.fullName}</h1>
        </div>

        <div className="flex flex-wrap justify-center gap-2">
          {profileConfig.contact.languages.map((lang) => (
            <div key={lang} className="bg-primary/5 rounded-full px-3 py-1">
              <p className="text-xs font-semibold text-primary">
                {lang}
              </p>
            </div>
          ))}
        </div>
      </div>

      <div className="text-sm mt-1">
        <div className="flex flex-col gap-2.5">
          <ContactItem
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path></svg>}
            text={profileConfig.contact.phone}
          />
          <ContactItem
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><rect width="20" height="16" x="2" y="4" rx="2"></rect><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"></path></svg>}
            text={profileConfig.contact.email}
            href={`mailto:${profileConfig.contact.email}`}
          />
          <ContactItem
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="m22 2-7 20-4-9-9-4Z"></path><path d="M22 2 11 13"></path></svg>}
            text={profileConfig.contact.telegram}
            href={`https://t.me/${profileConfig.contact.telegram.replace('@', '')}`}
          />
          <ContactItem
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"></path><rect width="4" height="12" x="2" y="9"></rect><circle cx="4" cy="4" r="2"></circle></svg>}
            text={profileConfig.contact.linkedin}
            href={`https://linkedin.com/in/${profileConfig.contact.linkedin}`}
          />
          <ContactItem
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4"></path><path d="M9 18c-4.51 2-5-2-7-2"></path></svg>}
            text={profileConfig.contact.github}
            href={`https://github.com/${profileConfig.contact.github}`}
          />

          <ContactItem
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"></path><circle cx="12" cy="10" r="3"></circle></svg>}
            text={
              <div className="flex flex-col gap-1 mt-0.5">
                <span>{profileConfig.contact.location_city}</span>
                <span className="inline-flex items-center gap-1.5 font-bold text-emerald-600 dark:text-emerald-400">
                  <span className="relative flex h-2 w-2">
                    <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
                    <span className="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span>
                  </span>
                  {profileConfig.contact.relocation}
                </span>
              </div>
            }
          />
        </div>
      </div>

      <div className="mt-5">
        <h3 className="text-xs font-bold text-muted-foreground uppercase tracking-widest mb-3 pl-2">
          {t('sections.experience', 'Work Experience')}
        </h3>
        <div className="flex flex-col gap-1.5 ml-1">
          {experience.map((exp, idx) => {
            const isActive = activeIndex === idx;
            return (
              <button
                key={idx}
                onClick={() => {
                  if (isActive) {
                    const event = new CustomEvent('open-experience', { detail: { index: -1 } });
                    window.dispatchEvent(event);
                  } else {
                    const event = new CustomEvent('open-experience', { detail: { index: idx, expandSubProjects: true } });
                    window.dispatchEvent(event);
                  }
                }}
                className="group flex items-center gap-3 w-full text-left cursor-pointer py-1"
              >
                <div
                  className={`shrink-0 h-[2px] transition-all duration-300 ease-out rounded-full ${isActive
                    ? 'w-6 bg-[var(--color-primary)]'
                    : 'w-3 bg-[var(--muted)] group-hover:w-6 group-hover:bg-[var(--color-primary)]'
                    }`}
                />
                <span
                  className={`text-sm font-medium transition-colors truncate ${isActive
                    ? 'text-[var(--color-primary)]'
                    : 'text-[var(--muted)] group-hover:text-[var(--color-primary)]'
                    }`}
                >
                  {exp.company}
                </span>
              </button>
            );
          })}
        </div>
      </div>
    </aside>
  );
}

function ContactItem({ icon, text, href }: { icon: React.ReactNode; text: React.ReactNode; href?: string }) {
  if (!text) return null;

  if (href) {
    return (
      <a
        href={href}
        target="_blank"
        rel="noopener noreferrer"
        className="flex items-start gap-3.5 text-muted-foreground hover:text-primary transition-colors group"
      >
        <span className="text-primary group-hover:text-primary transition-colors mt-0.5 shrink-0">{icon}</span>
        <span className="break-words leading-relaxed">{text}</span>
      </a>
    );
  }

  return (
    <div className="flex items-start gap-3.5 text-muted-foreground">
      <span className="text-primary mt-0.5 shrink-0">{icon}</span>
      <span className="break-words leading-relaxed">{text}</span>
    </div>
  );
}

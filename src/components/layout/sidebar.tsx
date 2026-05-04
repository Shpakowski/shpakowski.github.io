import { useTranslation } from 'react-i18next';
import { ContactItem } from '@/components/ui';
import { profileConfig } from '@/config';
import { getContactsConfig } from '@/config';
import { Avatar } from '@/components/ui';
import { useLayout } from '@/context';
import avatarImage from '@/assets/avatar.JPG';

export function Sidebar() {
  const { t } = useTranslation();
  const experience = t('experience', { returnObjects: true }) as Array<Record<string, string>>;
  const { activeExperienceIndex, openExperience } = useLayout();

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
          {getContactsConfig().map((contact) => {
            const Icon = contact.icon;
            return (
              <ContactItem
                key={contact.id}
                icon={<Icon size={18} />}
                text={contact.text}
                href={contact.href}
              />
            );
          })}
        </div>
      </div>

      <div className="mt-5">
        <h3 className="text-xs font-bold text-muted-foreground uppercase tracking-widest mb-3 pl-2">
          {t('sections.experience', 'Work Experience')}
        </h3>
        <div className="flex flex-col gap-1.5 ml-1">
          {experience.map((exp, idx) => {
            const isActive = activeExperienceIndex === idx;
            return (
              <button
                key={idx}
                onClick={() => {
                  if (isActive) {
                    openExperience(null);
                  } else {
                    openExperience(idx, true);
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

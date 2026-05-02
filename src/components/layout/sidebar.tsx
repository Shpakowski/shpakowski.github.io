
import { useTranslation } from 'react-i18next';
import { Avatar } from '../ui/avatar';
import { ThemeToggle } from '../ui/theme-toggle';
import { LanguageSwitch } from '../ui/language-switch';


import avatarImage from '../../assets/avatar.JPG';

interface SidebarProps {
  // We'll pass the profile data here or fetch it via i18n
  // Let's use i18n directly inside for now
}

export function Sidebar({}: SidebarProps) {
  const { t } = useTranslation();

  return (
    <aside className="flex flex-col gap-6 p-6 md:p-8 lg:sticky lg:top-0 lg:h-screen lg:w-80 xl:w-96 lg:overflow-y-auto bg-card/50">
      <div className="flex items-center justify-between">
        <ThemeToggle />
        <LanguageSwitch />
      </div>

      <div className="flex flex-col items-center text-center gap-4">
        <Avatar 
          src={avatarImage} 
          alt={t('profile.fullName')}
          className="h-32 w-32 shadow-md ring-2 ring-primary/20"
        />
        <div>
          <h1 className="text-2xl font-bold tracking-tight">{t('profile.fullName')}</h1>
          <p className="text-sm text-muted-foreground mt-1 font-medium">{t('profile.headline')}</p>
        </div>
        
        {/* Languages moved higher */}
        <div className="bg-primary/5 rounded-full px-4 py-1.5 mt-1">
          <p className="text-xs font-semibold text-primary text-center">
            {t('profile.contact.languages')}
          </p>
        </div>
      </div>

      <div className="space-y-4 text-sm mt-4">
        <div className="flex flex-col gap-3.5">
          <ContactItem 
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path></svg>} 
            text={t('profile.contact.phone')} 
          />
          <ContactItem 
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><rect width="20" height="16" x="2" y="4" rx="2"></rect><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"></path></svg>} 
            text={t('profile.contact.email')} 
            href={`mailto:${t('profile.contact.email')}`} 
          />
          <ContactItem 
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="m22 2-7 20-4-9-9-4Z"></path><path d="M22 2 11 13"></path></svg>} 
            text={t('profile.contact.telegram')} 
            href={`https://t.me/${t('profile.contact.telegram').replace('@', '')}`} 
          />
          <ContactItem 
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"></path><rect width="4" height="12" x="2" y="9"></rect><circle cx="4" cy="4" r="2"></circle></svg>} 
            text={t('profile.contact.linkedin')} 
            href={`https://${t('profile.contact.linkedin')}`} 
          />
          <ContactItem 
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4"></path><path d="M9 18c-4.51 2-5-2-7-2"></path></svg>} 
            text={t('profile.contact.github')} 
            href={`https://${t('profile.contact.github')}`} 
          />

          <ContactItem 
            icon={<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"></path><circle cx="12" cy="10" r="3"></circle></svg>} 
            text={
              <div className="flex flex-col gap-1 mt-0.5">
                <span>{t('profile.contact.location_city')}</span>
                <span className="inline-flex items-center gap-1.5 font-bold text-emerald-600 dark:text-emerald-400">
                  <span className="relative flex h-2 w-2">
                    <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
                    <span className="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span>
                  </span>
                  {t('profile.contact.relocation')}
                </span>
              </div>
            } 
          />
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

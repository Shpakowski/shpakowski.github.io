
import { useTranslation } from 'react-i18next';
import { Avatar } from '../ui/avatar';
import { ThemeToggle } from '../ui/theme-toggle';
import { LanguageSwitch } from '../ui/language-switch';
import { Mail, Phone, MapPin, Globe, Code, Briefcase, Send } from 'lucide-react';

import avatarImage from '../../assets/avatar.JPG';

interface SidebarProps {
  // We'll pass the profile data here or fetch it via i18n
  // Let's use i18n directly inside for now
}

export function Sidebar({}: SidebarProps) {
  const { t } = useTranslation();

  return (
    <aside className="flex flex-col gap-6 p-6 md:p-8 lg:sticky lg:top-0 lg:h-screen lg:w-80 lg:overflow-y-auto lg:border-r border-border bg-card/50">
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
      </div>

      <div className="space-y-4 text-sm mt-4">
        <div className="flex flex-col gap-3">
          <ContactItem icon={<Phone className="h-4 w-4" />} text={t('profile.contact.phone')} />
          <ContactItem icon={<Mail className="h-4 w-4" />} text={t('profile.contact.email')} href={`mailto:${t('profile.contact.email')}`} />
          <ContactItem icon={<Send className="h-4 w-4" />} text={t('profile.contact.telegram')} href={`https://t.me/${t('profile.contact.telegram').replace('@', '')}`} />
          <ContactItem icon={<Briefcase className="h-4 w-4" />} text={t('profile.contact.linkedin')} href={`https://${t('profile.contact.linkedin')}`} />
          <ContactItem icon={<Code className="h-4 w-4" />} text={t('profile.contact.github')} href={`https://${t('profile.contact.github')}`} />
          <ContactItem icon={<Globe className="h-4 w-4" />} text={t('profile.contact.website')} href={`https://${t('profile.contact.website')}`} />
          <ContactItem icon={<MapPin className="h-4 w-4" />} text={t('profile.contact.location')} />
        </div>
      </div>

      <div className="mt-auto pt-6 border-t border-border">
        <p className="text-xs text-muted-foreground text-center">
          {t('profile.contact.languages')}
        </p>
      </div>
    </aside>
  );
}

function ContactItem({ icon, text, href }: { icon: React.ReactNode; text: string; href?: string }) {
  if (!text) return null;
  
  if (href) {
    return (
      <a 
        href={href} 
        target="_blank" 
        rel="noopener noreferrer"
        className="flex items-center gap-3 text-muted-foreground hover:text-primary transition-colors group"
      >
        <span className="text-primary/70 group-hover:text-primary transition-colors">{icon}</span>
        <span className="truncate">{text}</span>
      </a>
    );
  }

  return (
    <div className="flex items-center gap-3 text-muted-foreground">
      <span className="text-primary/70">{icon}</span>
      <span className="truncate">{text}</span>
    </div>
  );
}

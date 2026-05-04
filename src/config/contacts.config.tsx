import { Phone, Mail, Send, MapPin } from 'lucide-react';
import { GithubIcon, LinkedinIcon } from '@/components/ui';
import { profileConfig } from './profile.config';
import type { ElementType, ReactNode } from 'react';

// 🎓 React Trend: Декларативная конфигурация для UI списков вместо хардкода элементов.
// Легко расширять, изменять порядок и мапить без дублирования JSX.
export interface ContactConfig {
  id: string;
  icon: ElementType;
  text: ReactNode;
  href?: string;
}

export const getContactsConfig = (): ContactConfig[] => [
  {
    id: 'phone',
    icon: Phone,
    text: profileConfig.contact.phone,
  },
  {
    id: 'email',
    icon: Mail,
    text: profileConfig.contact.email,
    href: `mailto:${profileConfig.contact.email}`,
  },
  {
    id: 'telegram',
    icon: Send,
    text: profileConfig.contact.telegram,
    href: `https://t.me/${profileConfig.contact.telegram.replace('@', '')}`,
  },
  {
    id: 'linkedin',
    icon: LinkedinIcon,
    text: profileConfig.contact.linkedin,
    href: `https://linkedin.com/in/${profileConfig.contact.linkedin}`,
  },
  {
    id: 'github',
    icon: GithubIcon,
    text: profileConfig.contact.github,
    href: `https://github.com/${profileConfig.contact.github}`,
  },
  {
    id: 'location',
    icon: MapPin,
    text: (
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
    ),
  },
];

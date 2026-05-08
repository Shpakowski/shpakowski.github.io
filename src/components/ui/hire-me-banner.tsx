import { Settings } from 'lucide-react';
import { useTranslation } from 'react-i18next';
import { cn } from '@/utils';

interface HireMeBannerProps {
  className?: string;
}

export function HireMeBanner({ className }: HireMeBannerProps) {
  const { t } = useTranslation();

  return (
    <a
      href="mailto:v.shpakowski@gmail.com"
      className={cn(
        "group flex items-center justify-center gap-1 sm:gap-4 cursor-pointer py-1.5 px-1 sm:px-3 rounded-full hover:bg-rose-500/5 transition-colors",
        className
      )}
    >
      {/* Left Spinning Gear */}
      <Settings
        className="text-rose-500 w-5 h-5 sm:w-7 sm:h-7 shrink-0 animate-[spin_4s_linear_infinite] group-hover:[animation-play-state:paused] transition-[animation-play-state]"
        strokeWidth={1.5}
        style={{ filter: 'drop-shadow(0px 1.5px 0.5px rgba(0,0,0,0.5))' }}
      />

      <span className="hidden sm:inline-block text-lg md:text-xl font-light tracking-tight text-rose-500 whitespace-nowrap">
        {t('profile.target')}
      </span>
      {/* Mobile large text */}
      <span className="sm:hidden text-base leading-none font-light tracking-tight text-rose-500 whitespace-nowrap uppercase">
        {t('profile.available')}
      </span>

      {/* Right Spinning Gear */}
      <Settings
        className="text-rose-500 w-5 h-5 sm:w-7 sm:h-7 shrink-0 animate-[spin_4s_linear_infinite] group-hover:[animation-play-state:paused] transition-[animation-play-state]"
        strokeWidth={1.5}
        style={{ filter: 'drop-shadow(0px 1.5px 0.5px rgba(0,0,0,0.5))' }}
      />
    </a>
  );
}

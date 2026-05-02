
import { useTranslation } from 'react-i18next';
import { cn } from '../../utils/cn';

export function LanguageSwitch() {
  const { i18n } = useTranslation();
  const currentLanguage = i18n.language || window.localStorage.i18nextLng || 'en';


  return (
    <div className="relative inline-flex h-9 w-32 items-center justify-center rounded-full bg-secondary/50 p-1">
      {/* Sliding Background */}
      <div
        className={cn(
          "absolute left-1 top-1 bottom-1 w-[calc(50%-4px)] rounded-full bg-background shadow-sm border border-primary/70 transition-transform duration-300 ease-in-out",
          currentLanguage === 'ru' ? "translate-x-full" : "translate-x-0"
        )}
      />
      
      {/* English Button */}
      <button
        onClick={() => i18n.changeLanguage('en')}
        className={cn(
          "relative z-10 flex h-full flex-1 items-center justify-center gap-1.5 rounded-full text-xs font-semibold transition-colors",
          currentLanguage === 'en' ? "text-foreground" : "text-muted-foreground hover:text-foreground"
        )}
      >
        <span className="text-sm">🇺🇸</span> EN
      </button>

      {/* Russian Button */}
      <button
        onClick={() => i18n.changeLanguage('ru')}
        className={cn(
          "relative z-10 flex h-full flex-1 items-center justify-center gap-1.5 rounded-full text-xs font-semibold transition-colors",
          currentLanguage === 'ru' ? "text-foreground" : "text-muted-foreground hover:text-foreground"
        )}
      >
        <span className="text-sm">🇷🇺</span> RU
      </button>
    </div>
  );
}

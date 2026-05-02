
import { useTranslation } from 'react-i18next';
import { cn } from '../../utils/cn';

export function LanguageSwitch() {
  const { i18n } = useTranslation();
  const currentLanguage = i18n.language || window.localStorage.i18nextLng || 'en';

  const toggleLanguage = () => {
    const newLang = currentLanguage === 'en' ? 'ru' : 'en';
    i18n.changeLanguage(newLang);
  };

  return (
    <button
      onClick={toggleLanguage}
      className="relative inline-flex h-8 w-16 items-center rounded-full bg-secondary/50 p-1 transition-colors hover:bg-secondary focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:ring-offset-background"
      aria-label="Toggle language"
    >
      <span className="sr-only">Toggle language</span>
      
      {/* Background text labels */}
      <span className="absolute inset-0 flex items-center justify-between px-2 text-[10px] font-bold text-muted-foreground">
        <span>EN</span>
        <span>RU</span>
      </span>

      {/* Thumb */}
      <span
        className={cn(
          "relative z-10 flex h-6 w-6 transform items-center justify-center rounded-full bg-background text-sm shadow-sm transition-transform duration-300",
          currentLanguage === 'ru' ? "translate-x-8" : "translate-x-0"
        )}
      >
        {currentLanguage === 'en' ? '🇺🇸' : '🇷🇺'}
      </span>
    </button>
  );
}

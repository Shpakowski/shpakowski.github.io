
import { useTranslation } from 'react-i18next';
import { cn } from '../../utils/cn';

export function LanguageSwitch() {
  const { i18n } = useTranslation();
  const currentLanguage = i18n.language || window.localStorage.i18nextLng || 'en';

  const toggleLanguage = () => {
    i18n.changeLanguage(currentLanguage === 'en' ? 'ru' : 'en');
  };

  return (
    <button
      onClick={toggleLanguage}
      className="relative flex h-9 w-[70px] items-center rounded-full border-2 border-foreground shadow-inner bg-secondary/50 p-1 transition-colors hover:bg-secondary/70 focus:outline-none"
      aria-label="Toggle language"
    >
      {/* Background Labels */}
      <div className="absolute inset-0 flex items-center justify-between px-0.5 text-xs font-bold text-muted-foreground opacity-70 pointer-events-none">
        <span className="w-7 text-center">EN</span>
        <span className="w-7 text-center">RU</span>
      </div>

      {/* Sliding Knob */}
      <div
        className={cn(
          "absolute left-0.5 top-0.5 z-10 h-7 w-7 rounded-full bg-background shadow-md border-2 border-foreground transition-transform duration-300 ease-in-out flex items-center justify-center overflow-hidden",
          currentLanguage === 'ru' ? "translate-x-[34px]" : "translate-x-0"
        )}
      >
        {currentLanguage === 'en' ? (
          <img src="https://flagcdn.com/us.svg" alt="EN" className="h-full w-full object-cover" />
        ) : (
          <img src="https://flagcdn.com/ru.svg" alt="RU" className="h-full w-full object-cover" />
        )}
      </div>
    </button>
  );
}

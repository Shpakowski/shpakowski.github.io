import { AppLayout } from './components/layout/app-layout';
import { HeroSection } from './components/sections/hero-section';
import { SkillsSection } from './components/sections/skills-section';
import { ExperienceSection } from './components/sections/experience-section';
import { LayoutProvider } from './context/LayoutContext';
import { useTranslation } from 'react-i18next';
import './config/i18n';

// 🎓 React Trend: Для больших проектов тяжелые компоненты секций рекомендуется загружать лениво 
// через React.lazy() и <Suspense> (Code Splitting). Это радикально ускоряет загрузку первого экрана (FCP).
// Для сайта-визитки прямой импорт допустим (KISS), но при масштабировании это стоит внедрить первым.
function App() {
  const { t } = useTranslation();

  return (
    <LayoutProvider>
      <AppLayout>
        <div className="space-y-12">
        <h1 className="text-xl md:text-2xl lg:text-3xl font-light tracking-tight text-center mb-6">
          {t('profile.headline')}
        </h1>
        
        <HeroSection />
        <SkillsSection />
        <ExperienceSection />
        </div>
      </AppLayout>
    </LayoutProvider>
  );
}

export default App;

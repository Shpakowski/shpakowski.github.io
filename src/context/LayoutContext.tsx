import { createContext, useContext, useState } from 'react';
import type { ReactNode } from 'react';

interface LayoutContextType {
  activeExperienceIndex: number | null;
  forceExpandSignal: number;
  openExperience: (index: number | null, expandSubProjects?: boolean) => void;
}

const LayoutContext = createContext<LayoutContextType | undefined>(undefined);

// 🎓 React Trend: Использование Context API (или Zustand) для глобального состояния UI.
// Это заменяет грязный антипаттерн с window.dispatchEvent и делает архитектуру предсказуемой (SOLID: DIP).
export function LayoutProvider({ children }: { children: ReactNode }) {
  const [activeExperienceIndex, setActiveExperienceIndex] = useState<number | null>(null);
  const [forceExpandSignal, setForceExpandSignal] = useState(0);

  const openExperience = (index: number | null, expandSubProjects: boolean = false) => {
    setActiveExperienceIndex(index);
    if (expandSubProjects) {
      setForceExpandSignal(prev => prev + 1);
    } else {
      setForceExpandSignal(0);
    }
  };

  return (
    <LayoutContext.Provider value={{
      activeExperienceIndex,
      forceExpandSignal,
      openExperience
    }}>
      {children}
    </LayoutContext.Provider>
  );
}

// eslint-disable-next-line react-refresh/only-export-components
export function useLayout() {
  const context = useContext(LayoutContext);
  if (context === undefined) {
    throw new Error('useLayout must be used within a LayoutProvider');
  }
  return context;
}

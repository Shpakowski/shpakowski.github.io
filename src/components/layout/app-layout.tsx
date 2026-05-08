import { Sidebar } from './sidebar';
import { ThemeToggle, LanguageSwitch, HireMeBanner } from '@/components/ui';


interface AppLayoutProps {
  children: React.ReactNode;
}

export function AppLayout({ children }: AppLayoutProps) {


  return (
    <div className="min-h-screen overflow-x-hidden bg-background text-foreground transition-colors duration-300">
      {/* 40px Header */}
      <header className="flex flex-wrap sm:flex-nowrap min-h-[40px] w-full items-center justify-center sm:justify-between gap-3 sm:gap-0 bg-background px-2 sm:px-4 md:px-8 my-5">
        {/* Invisible spacer for absolute centering on larger screens */}
        <div className="flex-1 hidden sm:block"></div>
        
        <div className="flex justify-center shrink-0">
          <HireMeBanner />
        </div>
        
        <div className="flex sm:flex-1 justify-center sm:justify-end gap-2 sm:gap-4">
          <LanguageSwitch />
          <ThemeToggle />
        </div>
      </header>
      
      <div className="mx-auto max-w-[1200px] flex flex-col lg:flex-row relative">
        <Sidebar />
        <main className="flex-1 overflow-x-hidden">
          <div className="px-6 md:px-8 lg:px-10">
            {children}
          </div>
        </main>
      </div>
    </div>
  );
}

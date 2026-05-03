import { Sidebar } from './sidebar';
import { ThemeToggle } from '../ui/theme-toggle';
import { LanguageSwitch } from '../ui/language-switch';


interface AppLayoutProps {
  children: React.ReactNode;
}

export function AppLayout({ children }: AppLayoutProps) {


  return (
    <div className="min-h-screen bg-background text-foreground transition-colors duration-300">
      {/* 40px Header */}
      <header className="flex h-[40px] w-full items-center justify-end gap-4 bg-background px-6 md:px-8 my-5">
        <LanguageSwitch />
        <ThemeToggle />
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

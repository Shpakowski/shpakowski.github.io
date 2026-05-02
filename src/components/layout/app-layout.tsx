
import { Sidebar } from './sidebar';

interface AppLayoutProps {
  children: React.ReactNode;
}

export function AppLayout({ children }: AppLayoutProps) {
  return (
    <div className="flex min-h-screen flex-col lg:flex-row bg-background text-foreground">
      <Sidebar />
      <main className="flex-1 overflow-x-hidden">
        <div className="mx-auto max-w-4xl p-6 md:p-8 lg:p-12">
          {children}
        </div>
      </main>
    </div>
  );
}

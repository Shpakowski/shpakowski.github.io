import { useState, useEffect } from 'react';

// 🎓 React Trend: Вынесение бизнес-логики и побочных эффектов (интервалов) 
// в пользовательские хуки для чистоты компонентов (SOLID: Single Responsibility Principle).
export function useCountUp(target: number, delay: number, duration: number = 1200) {
  const [count, setCount] = useState(0);
  const [started, setStarted] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => setStarted(true), delay * 1000);
    return () => clearTimeout(timer);
  }, [delay]);

  useEffect(() => {
    if (!started) return;
    const steps = 30;
    const increment = target / steps;
    let current = 0;
    const interval = setInterval(() => {
      current += increment;
      if (current >= target) {
        setCount(target);
        clearInterval(interval);
      } else {
        setCount(Math.floor(current));
      }
    }, duration / steps);
    return () => clearInterval(interval);
  }, [started, target, duration]);

  return count;
}

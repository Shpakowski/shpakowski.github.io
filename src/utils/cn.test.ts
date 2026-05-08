import { describe, it, expect } from 'vitest';
import { cn } from './cn';

describe('cn utility', () => {
  it('merges tailwind classes without conflicts', () => {
    expect(cn('bg-red-500', 'bg-blue-500')).toBe('bg-blue-500');
  });

  it('handles conditional classes', () => {
    const isMobile = true;
    const isDesktop = false;
    expect(cn('p-4', isMobile && 'm-4', isDesktop && 'p-2')).toBe('p-4 m-4');
  });

  it('merges arrays and objects correctly', () => {
    expect(cn(['text-sm', 'font-bold'], { 'text-red-500': true })).toBe('text-sm font-bold text-red-500');
  });
});

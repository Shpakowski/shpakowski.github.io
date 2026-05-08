import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { renderHook, act } from '@testing-library/react';
import { useCountUp } from './useCountUp';

describe('useCountUp', () => {
  beforeEach(() => {
    vi.useFakeTimers();
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  it('starts at 0 and reaches target value', () => {
    const { result } = renderHook(() => useCountUp(100, 0, 1000));
    expect(result.current).toBe(0);

    // Initial timeout for start
    act(() => {
      vi.advanceTimersByTime(10);
    });

    act(() => {
      vi.advanceTimersByTime(500);
    });
    expect(result.current).toBeGreaterThan(0);
    expect(result.current).toBeLessThan(100);

    act(() => {
      vi.advanceTimersByTime(600);
    });
    expect(result.current).toBe(100);
  });

  it('respects delay before starting', () => {
    const { result } = renderHook(() => useCountUp(100, 1, 1000)); // 1 sec delay
    expect(result.current).toBe(0);

    act(() => {
      vi.advanceTimersByTime(500); // Wait 500ms (still delayed)
    });
    expect(result.current).toBe(0);

    act(() => {
      vi.advanceTimersByTime(600); // 1100ms total (past delay)
    });

    act(() => {
      vi.advanceTimersByTime(100); // Time for increment
    });
    expect(result.current).toBeGreaterThan(0);
  });
});

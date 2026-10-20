import '@testing-library/jest-dom/vitest';
import { vi } from 'vitest';

const store = new Map<string, string>();
const storageMock = new Proxy({} as Storage, {
  get(_target, prop: string) {
    if (prop === 'getItem') return (key: string) => store.get(key) ?? null;
    if (prop === 'setItem') return (key: string, val: string) => { store.set(key, String(val)); };
    if (prop === 'removeItem') return (key: string) => { store.delete(key); };
    if (prop === 'clear') return () => { store.clear(); };
    if (prop === 'length') return store.size;
    if (prop === 'key') return (idx: number) => Array.from(store.keys())[idx] ?? null;
    return store.get(prop);
  },
  set(_target, prop: string, value: unknown) {
    store.set(prop, String(value));
    return true;
  },
  has(_target, prop: string) {
    return store.has(prop);
  },
  deleteProperty(_target, prop: string) {
    return store.delete(prop);
  }
});

Object.defineProperty(globalThis, 'localStorage', {
  value: storageMock,
  writable: true,
  configurable: true,
});
Object.defineProperty(window, 'localStorage', {
  value: storageMock,
  writable: true,
  configurable: true,
});

Object.defineProperty(window, 'matchMedia', {
  writable: true,
  value: vi.fn().mockImplementation(query => ({
    matches: false,
    media: query,
    onchange: null,
    addListener: vi.fn(),
    removeListener: vi.fn(),
    addEventListener: vi.fn(),
    removeEventListener: vi.fn(),
    dispatchEvent: vi.fn(),
  })),
});

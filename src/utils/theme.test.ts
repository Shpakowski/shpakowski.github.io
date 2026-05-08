import { describe, it, expect } from 'vitest';
import { getThemeForSkill } from './theme';

describe('getThemeForSkill', () => {
  it('returns theme for exact match', () => {
    expect(typeof getThemeForSkill('react')).toBe('object');
  });

  it('returns theme for partial match', () => {
    expect(typeof getThemeForSkill('react native')).toBe('object');
  });

  it('returns theme from fallback map', () => {
    expect(typeof getThemeForSkill('go')).toBe('object');
    expect(typeof getThemeForSkill('vite')).toBe('object');
  });

  it('returns null for unknown skills', () => {
    expect(getThemeForSkill('unknown_technology')).toBeNull();
  });
});

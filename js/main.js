/**
 * Simple CV Website
 * Language and Theme Switching Only
 */

'use strict';

// =============================================================================
// Constants
// =============================================================================
const STORAGE_KEYS = {
    LANGUAGE: 'cv-language',
    THEME: 'cv-theme'
};

const THEMES = {
    LIGHT: 'light',
    DARK: 'dark'
};

// =============================================================================
// DOM Elements
// =============================================================================
const elements = {
    languageBtns: document.querySelectorAll('.language-switcher__btn'),
    themeToggle: document.querySelector('.theme-toggle'),
    themeIcon: document.querySelector('.theme-toggle__icon'),
    resumeSections: document.querySelectorAll('.resume')
};

// =============================================================================
// Language Switching
// =============================================================================

/**
 * Switch language
 * @param {string} lang - Language code (ru/en)
 */
function switchLanguage(lang) {
    // Update resume sections visibility
    elements.resumeSections.forEach(section => {
        if (section.dataset.lang === lang) {
            section.classList.add('resume--active');
        } else {
            section.classList.remove('resume--active');
        }
    });
    
    // Update language buttons
    elements.languageBtns.forEach(btn => {
        if (btn.dataset.lang === lang) {
            btn.classList.add('language-switcher__btn--active');
        } else {
            btn.classList.remove('language-switcher__btn--active');
        }
    });
    
    // Update HTML lang attribute
    document.documentElement.setAttribute('lang', lang);
    
    // Save to localStorage
    try {
        localStorage.setItem(STORAGE_KEYS.LANGUAGE, lang);
    } catch (e) {
        console.warn('Failed to save language preference');
    }
}

/**
 * Initialize language switcher
 */
function initLanguageSwitcher() {
    elements.languageBtns.forEach(btn => {
        btn.addEventListener('click', () => {
            const lang = btn.dataset.lang;
            switchLanguage(lang);
        });
    });
    
    // Load saved language or default to EN
    let savedLang = 'en';
    try {
        savedLang = localStorage.getItem(STORAGE_KEYS.LANGUAGE) || 'en';
    } catch (e) {
        console.warn('Failed to load language preference');
    }
    
    switchLanguage(savedLang);
}

// =============================================================================
// Theme Switching
// =============================================================================

/**
 * Get system theme preference
 * @returns {string} 'dark' or 'light'
 */
function getSystemTheme() {
    return window.matchMedia('(prefers-color-scheme: dark)').matches 
        ? THEMES.DARK 
        : THEMES.LIGHT;
}

/**
 * Apply theme
 * @param {string} theme - Theme name (light/dark)
 */
function applyTheme(theme) {
    document.documentElement.setAttribute('data-theme', theme);
    
    // Update icon
    if (elements.themeIcon) {
        elements.themeIcon.textContent = theme === THEMES.DARK ? '☀️' : '🌙';
    }
    
    // Save to localStorage
    try {
        localStorage.setItem(STORAGE_KEYS.THEME, theme);
    } catch (e) {
        console.warn('Failed to save theme preference');
    }
}

/**
 * Toggle theme
 */
function toggleTheme() {
    const currentTheme = document.documentElement.getAttribute('data-theme') || THEMES.LIGHT;
    const newTheme = currentTheme === THEMES.LIGHT ? THEMES.DARK : THEMES.LIGHT;
    applyTheme(newTheme);
}

/**
 * Initialize theme switcher
 */
function initThemeSwitcher() {
    // Load saved theme or use system preference
    let savedTheme = getSystemTheme();
    try {
        savedTheme = localStorage.getItem(STORAGE_KEYS.THEME) || getSystemTheme();
    } catch (e) {
        console.warn('Failed to load theme preference');
    }
    
    applyTheme(savedTheme);
    
    // Add click listener
    if (elements.themeToggle) {
        elements.themeToggle.addEventListener('click', toggleTheme);
    }
    
    // Listen for system theme changes
    window.matchMedia('(prefers-color-scheme: dark)')
        .addEventListener('change', (e) => {
            // Only apply system theme if user hasn't manually set one
            try {
                if (!localStorage.getItem(STORAGE_KEYS.THEME)) {
                    applyTheme(e.matches ? THEMES.DARK : THEMES.LIGHT);
                }
            } catch (err) {
                console.warn('Failed to handle system theme change');
            }
        });
}

// =============================================================================
// Initialization
// =============================================================================

/**
 * Initialize the application
 */
function init() {
    initLanguageSwitcher();
    initThemeSwitcher();
    console.log('✅ CV Website initialized');
}

// Wait for DOM to be ready
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
} else {
    init();
}

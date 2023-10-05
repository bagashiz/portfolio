/*!
 * Minimal theme switcher
 *
 * Pico.css - https://picocss.com
 * Copyright 2019-2023 - Licensed under MIT
 */

/**
 * Minimal theme switcher
 *
 * @namespace
 * @typedef {Object} ThemeSwitcher
 * @property {string} _scheme - The current color scheme ("auto", "light", or "dark").
 * @property {string} menuTarget - The selector for the menu element that contains theme switchers.
 * @property {string} buttonsTarget - The selector for theme switcher buttons.
 * @property {string} buttonAttribute - The attribute name used for theme switcher buttons.
 * @property {string} rootAttribute - The attribute name used for the root HTML element to store the selected theme.
 * @property {string} localStorageKey - The key used to store the preferred color scheme in local storage.
 */
export const ThemeSwitcher = {
	// Config
	_scheme: 'auto',
	menuTarget: "details[role='list']",
	buttonsTarget: 'a[data-theme-switcher]',
	buttonAttribute: 'data-theme-switcher',
	rootAttribute: 'data-theme',
	localStorageKey: 'picoPreferredColorScheme',

	/**
	 * Initialize the theme switcher.
	 *
	 * @function
	 * @memberof ThemeSwitcher
	 */
	init() {
		this.scheme = this.schemeFromLocalStorage || this.preferredColorScheme;
		this.initSwitchers();
	},

	/**
	 * Get the color scheme from local storage or use the preferred color scheme.
	 *
	 * @function
	 * @memberof ThemeSwitcher
	 * @returns {string|null} The color scheme ("light", "dark", or null).
	 */
	get schemeFromLocalStorage() {
		if (typeof window.localStorage !== 'undefined') {
			if (window.localStorage.getItem(this.localStorageKey) !== null) {
				return window.localStorage.getItem(this.localStorageKey);
			}
		}
		return this._scheme;
	},

	/**
	 * Get the preferred color scheme based on user preferences.
	 *
	 * @function
	 * @memberof ThemeSwitcher
	 * @returns {string} The preferred color scheme ("light" or "dark").
	 */
	get preferredColorScheme() {
		return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
	},

	/**
	 * Initialize the theme switcher buttons and their click events.
	 *
	 * @function
	 * @memberof ThemeSwitcher
	 */
	initSwitchers() {
		const buttons = document.querySelectorAll(this.buttonsTarget);
		buttons.forEach((button) => {
			button.addEventListener(
				'click',
				(event) => {
					event.preventDefault();
					// Set scheme
					this.scheme = button.getAttribute(this.buttonAttribute) || 'auto';
					// Close dropdown
					document.querySelector(this.menuTarget)?.removeAttribute('open');
				},
				false
			);
		});
	},

	/**
	 * Set the selected color scheme and update the UI.
	 *
	 * @function
	 * @memberof ThemeSwitcher
	 * @param {string} scheme - The color scheme to set ("auto", "light", or "dark").
	 */
	set scheme(scheme) {
		if (scheme == 'auto') {
			this.preferredColorScheme == 'dark' ? (this._scheme = 'dark') : (this._scheme = 'light');
		} else if (scheme == 'dark' || scheme == 'light') {
			this._scheme = scheme;
		}
		this.applyScheme();
		this.schemeToLocalStorage();
	},

	/**
	 * Get the current color scheme.
	 *
	 * @function
	 * @memberof ThemeSwitcher
	 * @returns {string} The current color scheme ("auto", "light", or "dark").
	 */
	get scheme() {
		return this._scheme;
	},

	/**
	 * Apply the selected color scheme to the HTML root element.
	 *
	 * @function
	 * @memberof ThemeSwitcher
	 */
	applyScheme() {
		document.querySelector('html')?.setAttribute(this.rootAttribute, this.scheme);
	},

	/**
	 * Store the selected color scheme in local storage.
	 *
	 * @function
	 * @memberof ThemeSwitcher
	 */
	schemeToLocalStorage() {
		if (typeof window.localStorage !== 'undefined') {
			window.localStorage.setItem(this.localStorageKey, this.scheme);
		}
	}
};

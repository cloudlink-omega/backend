/*!
 * Color mode toggler with OS change detection
 */

(() => {
	'use strict';

	// Function to update theme based on localStorage or system preference
	const updateTheme = () => {
		if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
			document.documentElement.classList.add('dark');
		} else {
			document.documentElement.classList.remove('dark');
		}
	};

	// Initialize theme on load
	updateTheme();

	// Listen for OS color scheme changes
	window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
		updateTheme();
	});
})();
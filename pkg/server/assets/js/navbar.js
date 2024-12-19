const mobileMenuButton = document.getElementById('mobile-menu-button');
const mobileMenu = document.getElementById('mobile-menu');

mobileMenuButton.addEventListener('click', function () {
    const isExpanded = this.getAttribute('aria-expanded') === 'true';

    this.setAttribute('aria-expanded', !isExpanded);
    mobileMenu.classList.toggle('hidden');

    // Toggle icons
    const icons = this.getElementsByTagName('svg');
    icons[0].classList.toggle('hidden');
    icons[1].classList.toggle('hidden');
});

const desktopDropdownButton = document.getElementById('desktop-dropdown-button');
const desktopDropdownMenu = document.getElementById('desktop-dropdown-menu');

desktopDropdownButton.addEventListener('click', function (e) {
    e.stopPropagation();
    desktopDropdownMenu.classList.toggle('hidden');
});

const mobileDropdownButton = document.getElementById('mobile-dropdown-button');
const mobileDropdownMenu = document.getElementById('mobile-dropdown-menu');

mobileDropdownButton.addEventListener('click', function () {
    mobileDropdownMenu.classList.toggle('hidden');
});

document.addEventListener('click', function (e) {
    if (!desktopDropdownButton.contains(e.target)) {
        desktopDropdownMenu.classList.add('hidden');
    }
});

document.addEventListener('keydown', function (e) {
    if (e.key === 'Escape') {
        desktopDropdownMenu.classList.add('hidden');
        mobileDropdownMenu.classList.add('hidden');
    }
});
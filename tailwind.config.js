/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/*.html"],
    theme: {
        screens: {
            mobile: "360px",
            tablet: "601px",
            desktop: "1280px",
        },
    },
    plugins: [],
};

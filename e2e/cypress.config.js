const { defineConfig } = require("cypress");

module.exports = defineConfig({
  e2e: {
    experimentalSessionAndOrigin: true,
    viewportWidth: 1800,
    viewportHeight: 1000,
    baseUrl: "http://localhost:3000",
    supportFile: false,
    videoUploadOnPasses: false,
    setupNodeEvents(on, config) {},
  },
});

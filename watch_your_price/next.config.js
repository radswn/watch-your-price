const withPWA = require('next-pwa');

module.exports = withPWA({
    pwa: {
      disable: true,
      register: true
    },
    async redirects() {
      return [{
        source: '/',
        destination: '/search',
        permanent: true
      }];
    }
  })
module.exports = {
  testEnvironment: 'jsdom',
  testEnvironmentOptions: {
    customExportConditions: ['node', 'node-addons'],
  },
  cacheDirectory: '<rootDir>/tmp/cache/jest',

  preset: 'ts-jest',
  moduleFileExtensions: ['js', 'ts', 'vue'],
  moduleNameMapper: {
    '^@/(.*)$': '<rootDir>/src/$1',
    '\\.(css|less|scss|sass)$': 'identity-obj-proxy',
    '\\.(png|jpeg|jpg|gif)$': 'identity-obj-proxy',
    '\\.(ttf|woff|woff2)$': 'identity-obj-proxy',
    'iconfont\\.js$': 'identity-obj-proxy',
  },
  snapshotSerializers: ['<rootDir>/node_modules/jest-serializer-vue'],
  watchPlugins: ['jest-watch-typeahead/filename', 'jest-watch-typeahead/testname'],

  testRegex: '(/tests/.*|(\\.|/)(test|spec))\\.(js|ts)$',
  transformIgnorePatterns: ['/node_modules/'],
  transform: {
    '^.+\\.(js|ts)$': '@sucrase/jest-plugin',
    '^.+\\.(vue)$': '@vue/vue3-jest',
  },

  coverageDirectory: '<rootDir>/tmp/coverage',
  coverageReporters: ['html', 'lcov', 'text'],
  coverageProvider: 'v8',
  collectCoverageFrom: ['<rootDir>/src/**/*.{js,ts,vue}'],
  coverageThreshold: {
    global: {
      branches: 40,
      functions: 80,
      lines: 90,
      statements: 80,
    },
  },
}

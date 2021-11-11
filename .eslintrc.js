module.exports = {
  env: {
    browser: true,
    es2021: true,
    node: true,
  },
  extends: [
    'plugin:vue/essential',
    'airbnb-base',
  ],
  parserOptions: {
    ecmaVersion: 13,
    parser: '@typescript-eslint/parser',
    sourceType: 'module',
  },
  plugins: [
    'vue',
    '@typescript-eslint',
  ],
  rules: {
    'import/no-unresolved': 'off',
    'import/extensions': 'off',
    'import/no-extraneous-dependencies': 'off',
    'no-param-reassign': 'off',
    'vue/no-v-model-argument': 'off',
    camelcase: 'off',
    'import/prefer-default-export': 'off',
    'vue/no-multiple-template-root': 'off',
    'guard-for-in': 'off',
    'no-useless-return': 'off',
    'no-restricted-syntax': 'off',
    'no-unused-vars': 'off',
  },
};
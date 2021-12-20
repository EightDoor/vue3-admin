declare module '*.vue' {
  import { defineComponent } from 'vue';

  const component: ReturnType<typeof defineComponent>;
  export default component;
}

interface ImportMeta {
    readonly globEager: any
}

(window as any).global = window;

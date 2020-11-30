import * as path from "path";

const pathAliasMap = {
    '@/': '/src/',
}
export default  {
    resolvers: [
        {
            alias(path: string) {
                for (const [slug, res] of Object.entries(pathAliasMap)) {
                    if (path.startsWith(slug)) {
                        return path.replace(slug, res)
                    }
                }
            },
            proxy: {
                "api/":{
                    target: 'http://xxx.com',
                    changeOrigin: true,
                    rewrite: path => path.replace(/^\/api/, '')
                }
            },
        },
    ],
}

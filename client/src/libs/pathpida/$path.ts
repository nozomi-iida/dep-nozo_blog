export const pagesPath = {
  "articles": {
    _id: (id: string | number) => ({
      $url: (url?: { hash?: string }) => ({ pathname: '/articles/[id]' as const, query: { id }, hash: url?.hash })
    })
  },
  "search": {
    $url: (url?: { hash?: string }) => ({ pathname: '/search' as const, hash: url?.hash })
  },
  "topics": {
    _topic: (topic: string | number) => ({
      $url: (url?: { hash?: string }) => ({ pathname: '/topics/[topic]' as const, query: { topic }, hash: url?.hash })
    })
  },
  $url: (url?: { hash?: string }) => ({ pathname: '/' as const, hash: url?.hash })
}

export type PagesPath = typeof pagesPath

export const staticPath = {
  applause_png: '/applause.png',
  blog_test_jpeg: '/blog_test.jpeg',
  favicon_ico: '/favicon.ico',
  vercel_svg: '/vercel.svg'
} as const

export type StaticPath = typeof staticPath

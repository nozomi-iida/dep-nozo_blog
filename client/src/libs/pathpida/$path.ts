import type { OptionalQuery as OptionalQuery0 } from '../../pages/search/index.page'

export const pagesPath = {
  "admin": {
    "managements": {
      "articles": {
        _id: (id: string | number) => ({
          "edit": {
            $url: (url?: { hash?: string }) => ({ pathname: '/admin/managements/articles/[id]/edit' as const, query: { id }, hash: url?.hash })
          }
        }),
        "create": {
          $url: (url?: { hash?: string }) => ({ pathname: '/admin/managements/articles/create' as const, hash: url?.hash })
        },
        $url: (url?: { hash?: string }) => ({ pathname: '/admin/managements/articles' as const, hash: url?.hash })
      },
      "topics": {
        _id: (id: string | number) => ({
          "edit": {
            $url: (url?: { hash?: string }) => ({ pathname: '/admin/managements/topics/[id]/edit' as const, query: { id }, hash: url?.hash })
          }
        }),
        "create": {
          $url: (url?: { hash?: string }) => ({ pathname: '/admin/managements/topics/create' as const, hash: url?.hash })
        },
        $url: (url?: { hash?: string }) => ({ pathname: '/admin/managements/topics' as const, hash: url?.hash })
      }
    },
    "sign_in": {
      $url: (url?: { hash?: string }) => ({ pathname: '/admin/sign-in' as const, hash: url?.hash })
    }
  },
  "articles": {
    _id: (id: string | number) => ({
      $url: (url?: { hash?: string }) => ({ pathname: '/articles/[id]' as const, query: { id }, hash: url?.hash })
    })
  },
  "search": {
    $url: (url?: { query?: OptionalQuery0, hash?: string }) => ({ pathname: '/search' as const, query: url?.query, hash: url?.hash })
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
  blog_test_jpeg: '/blog_test.jpeg',
  clap_png: '/clap.png',
  clap_dark_png: '/clap_dark.png',
  favicon_ico: '/favicon.ico',
  nozomi_private_png: '/nozomi_private.png',
  nozomi_work_jpeg: '/nozomi_work.jpeg',
  vercel_svg: '/vercel.svg'
} as const

export type StaticPath = typeof staticPath

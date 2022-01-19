import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'

/* Layout */
import Layout from '@/layout/index.vue'

/* Router modules */
import componentsRouter from './modules/components'
import chartsRouter from './modules/charts'
import tableRouter from './modules/table'
import nestedRouter from './modules/nested'

Vue.use(VueRouter)

/*
  Note: sub-menu only appear when children.length>=1
  Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
*/

/*
  name:'router-name'             the name field is required when using <keep-alive>, it should also match its component's name property
                                 detail see : https://vuejs.org/v2/guide/components-dynamic-async.html#keep-alive-with-Dynamic-Components
  redirect:                      if set to 'noredirect', no redirect action will be trigger when clicking the breadcrumb
  meta: {
    roles: ['admin', 'editor']   will control the page roles (allow setting multiple roles)
    title: 'title'               the name showed in subMenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon showed in the sidebar
    hidden: true                 if true, this route will not show in the sidebar (default is false)
    alwaysShow: true             if true, will always show the root menu (default is false)
                                 if false, hide the root menu when has less or equal than one children route
    breadcrumb: false            if false, the item will be hidden in breadcrumb (default is true)
    noCache: true                if true, the page will not be cached (default is false)
    affix: true                  if true, the tag will affix in the tags-view
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
*/

/**
  ConstantRoutes
  a base page that does not have permission requirements
  all roles can be accessed
*/
export const constantRoutes: RouteConfig[] = [
  {
    path: '/redirect',
    component: Layout,
    meta: { hidden: true },
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import(/* webpackChunkName: "redirect" */ '@/views/redirect/index.vue')
      }
    ]
  },
  {
    path: '/login',
    component: () => import(/* webpackChunkName: "login" */ '@/views/login/index.vue'),
    meta: { hidden: true }
  },
  {
    path: '/auth-redirect',
    component: () => import(/* webpackChunkName: "auth-redirect" */ '@/views/login/auth-redirect.vue'),
    meta: { hidden: true }
  },
  {
    path: '/404',
    component: () => import(/* webpackChunkName: "404" */ '@/views/error-page/404.vue'),
    meta: { hidden: true }
  },
  {
    path: '/401',
    component: () => import(/* webpackChunkName: "401" */ '@/views/error-page/401.vue'),
    meta: { hidden: true }
  },
  {
    path: '/',
    component: Layout,
    meta: { name: 'products' },
    children: [
      {
        path: '',
        component: () => import('@/views/products/index.vue'),
        name: 'Products',
        meta: {
          title: 'products',
          icon: 'token',
          noCache: true
        }
      }
    ]
  },
  {
    path: '/products',
    component: Layout,
    meta: { hidden: true },
    children: [
      {
        path: ':pair([a-zA-Z-]+)',
        component: () => import('@/views/products/detail.vue'),
        name: 'ProductsDetail',
        meta: {
          title: 'products',
          noCache: true
        }
      }
    ]
  },
  {
    path: '/portfolio',
    component: Layout,
    meta: { name: 'portfolio', hidden: true },
    // redirect: 'noredirect',
    children: [
      {
        path: '',
        component: () => import('@/views/portfolio/index.vue'),
        name: 'Portfolio',
        meta: {
          title: 'portfolio',
          noCache: true
        }
      }
    ]
  },
  //   {
  //     path: '/products/pair',
  //     component: Layout,
  //     meta: { hidden: true },
  //     children: [
  //       {
  //         path: '',
  //         component: () => import('@/views/products/detail.vue'),
  //         name: 'ProductsDetail',
  //         meta: {
  //           title: 'productsDetail',
  //           noCache: true
  //         }
  //       }
  //     ]
  //   },
  {
    name: 'newpair',
    path: '/newpair',
    meta: {
      name: 'newpair',
      roles: ['admin', 'user']
    },
    component: Layout,
    children: [
      {
        path: 'index',
        component: () => import('@/views/pairs/newpair.vue'),
        name: 'newpair',
        meta: { title: 'add', icon: 'add', affix: true, roles: ['admin', 'user'] }
      }
    ]
  },
  {
    path: '/help',
    component: Layout,
    meta: { align: 'bottom', roles: ['admin', 'user'] },
    redirect: 'noredirect',
    children: [
      {
        path: 'index',
        component: () => import('@/views/help/index.vue'),
        name: 'Help',
        meta: {
          title: 'help',
          icon: 'help',
          affix: true,
          roles: ['admin', 'user']
        }
      }
    ]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
*/
export const asyncRoutes: RouteConfig[] = [
  {
    path: '*',
    redirect: '/404',
    meta: { hidden: true }
  },
  {
    path: '/template',
    component: Layout,
    meta: {
      name: 'template',
      roles: ['admin']
    },
    // redirect: 'noredirect',
    children: [
      {
        path: '',
        component: () => import('@/views/template/index.vue'),
        name: 'Template',
        meta: {
          title: 'template',
          icon: 'template',
          noCache: true,
          roles: ['admin']
        }
      }
    ]
  },
  {
    path: '/pairs',
    component: Layout,
    meta: {
      name: 'pairs',
      roles: ['admin']
    },
    // redirect: 'noredirect',
    children: [
      {
        path: '',
        component: () => import('@/views/pairs/index.vue'),
        name: 'Pairs',
        meta: {
          title: 'pairs',
          icon: 'token',
          noCache: true,
          roles: ['admin']
        }
      }
    ]
  },
  {
    path: '/setting',
    component: Layout,
    meta: {
      name: 'setting',
      roles: ['admin']
    },
    // redirect: 'noredirect',
    children: [
      {
        path: '',
        component: () => import('@/views/settings/index.vue'),
        name: 'Setting',
        meta: {
          title: 'setting',
          icon: 'setting',
          noCache: true,
          roles: ['admin']
        }
      }
    ]
  },
  {
    name: 'more',
    path: '/more',
    meta: {
      name: 'newpair',
      roles: ['admin', 'user']
    },
    component: Layout,
    redirect: 'noredirect',
    children: [
      {
        path: 'index',
        component: () => import('@/views/more/index.vue'),
        name: 'More',
        meta: {
          title: 'down',
          icon: 'down',
          affix: true,
          roles: ['admin', 'user']
        }
      }
    ]
  }
]

const createRouter = () => new VueRouter({
  mode: 'history', // Disabled due to Github Pages doesn't support this, enable this if you need.
  scrollBehavior: (to, from, savedPosition) => {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  },
  base: process.env.BASE_URL,
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter();
  (router as any).matcher = (newRouter as any).matcher // reset router
}

export default router

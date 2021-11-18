import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'
import { constantRoutes } from '@/router'
import Layout from '@/layout'
import { createPairs } from '@/common/PairList'

NProgress.configure({ showSpinner: false }) // NProgress Configuration

const whiteList = ['/login'] // no redirect whitelist

router.beforeEach(async (to, from, next) => {
  // start progress bar
  NProgress.start()

  // set page title
  document.title = getPageTitle(to.meta.title)

  // determine whether the user has logged in
  const hasToken = getToken()

  if (hasToken) {
    if (to.path === '/login') {
      // if is logged in, redirect to the home page
      next({ path: '/' })
      NProgress.done()
    } else {
      const hasGetUserInfo = store.getters.name
      if (hasGetUserInfo) {
        next()
      } else {
        try {
          // get user info
          await store.dispatch('user/getInfo')
          // const tempAsyncRoutes = Object.assign([], asyncRoutes)
          var pairObj = createPairs()
          var pairList = await pairObj.getPairList()
          // await store.dispatch('getPairsList', {}).then(async data => {
          //   if (data === null) {
          //     pairList = await pairObj.getPairList()
          //     store.dispatch('setPairsList', pairList)
          //   } else {
          //     pairList = data
          //   }
          // })
          store.state.pairsList = pairList
          console.log('pair name=', pairObj.name)
          console.log('pairs length=', pairList.size())
          var pairsRouter = [{
            path: 'index',
            name: 'Pairs',
            component: () => import('@/views/pairs/index'),
            // props: router => ({ pairindex: router.query.pi }),
            meta: { title: pairList.get(0).token0.symbol + ' / ' + pairList.get(0).token1.symbol, icon: 'table' }
          }]
          var routerItem
          // for (var p = 0; p < pairList.size(); p++) {
          routerItem = {
            path: 'index',
            name: 'Pairs',
            component: () => import('@/views/pairs/index'),
            // props: router => ({ pairindex: router.query.pi }),
            meta: { title: pairList.get(0).token0.symbol + ' / ' + pairList.get(0).token1.symbol, icon: 'table' }
          }
          // pairsRouter.push(routerItem)
          // }

          for (var i = 0; i < constantRoutes.length; i++) {
            if (constantRoutes[i].meta !== undefined && constantRoutes[i].meta.title === 'Pairs') {
              console.log('pair=', constantRoutes[i].redirect)
              constantRoutes[i].children = pairsRouter
            }
          }
          // const menu = {
          //   path: '/setting1',
          //   component: Layout,
          //   children: [{
          //     path: 'index',
          //     name: 'Test',
          //     component: () => import('@/views/setting/index'),
          //     meta: { title: 'Setting12', icon: 'el-icon-setting' }
          //   }]
          // }
          // const newRouters = constantRoutes.concat(menu)
          // console.log('newRouters=', newRouters)
          router.options.routes = constantRoutes
          router.addRoutes(constantRoutes)
          next({ ...to, replace: true })
        } catch (error) {
          // remove token and go to login page to re-login
          await store.dispatch('user/resetToken')
          Message.error(error || 'Has Error')
          next(`/login?redirect=${to.path}`)
          NProgress.done()
        }
      }
    }
  } else {
    /* has no token*/

    if (whiteList.indexOf(to.path) !== -1) {
      // in the free login whitelist, go directly
      next()
    } else {
      // other pages that do not have permission to access are redirected to the login page.
      next(`/login?redirect=${to.path}`)
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})

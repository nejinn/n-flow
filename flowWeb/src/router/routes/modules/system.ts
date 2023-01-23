import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';
import { RoleEnum } from '/@/enums/roleEnum';

const system: AppRouteModule = {
  path: '/system',
  name: 'System',
  component: LAYOUT,
  redirect: '/system/account',
  meta: {
    orderNo: 2000,
    icon: 'ion:settings-outline',
    title: t('routes.demo.system.moduleName'),
    roles: [RoleEnum.ADMINSYSTEM],
  },
  children: [
    {
      path: 'account',
      name: 'AccountManagement',
      meta: {
        title: t('routes.system.account'),
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/system/account/index.vue'),
    },
    // {
    //   path: 'account_detail/:id',
    //   name: 'AccountDetail',
    //   meta: {
    //     hideMenu: true,
    //     title: t('routes.demo.system.account_detail'),
    //     ignoreKeepAlive: true,
    //     showMenu: false,
    //     currentActiveMenu: '/system/account',
    //   },
    //   component: () => import('../../../views/system/account/AccountDetail.vue'),
    // },
    {
      path: 'role',
      name: 'RoleManagement',
      meta: {
        title: t('routes.system.role'),
        ignoreKeepAlive: true,
        roles: [RoleEnum.ADMINSYSTEMROLEGETROLELIST],
      },
      component: () => import('/@/views/system/role/index.vue'),
    },

    {
      path: 'permit',
      name: 'PermitManagement',
      meta: {
        title: t('routes.system.permit'),
        ignoreKeepAlive: true,
        roles: [RoleEnum.ADMINSYSTEMPERMIT],
      },
      component: () => import('/@/views/system/permit/index.vue'),
    },
    {
      path: 'dept',
      name: 'DeptManagement',
      meta: {
        title: t('routes.system.dept'),
        ignoreKeepAlive: true,
        roles: [RoleEnum.ADMINSYSTEMDEPT],
      },
      component: () => import('/@/views/system/dept/index.vue'),
    },
  ],
};

export default system;

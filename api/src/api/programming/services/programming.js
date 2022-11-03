'use strict';

/**
 * programming service
 */

const { createCoreService } = require('@strapi/strapi').factories;

module.exports = createCoreService('api::programming.programming');

/* tslint:disable */
/* eslint-disable */
/**
 * Daytona Runner API
 * Daytona Runner API
 *
 * The version of the OpenAPI document: v0.0.0-dev
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

// May contain unused imports in some cases
// @ts-ignore
import type { RegistryDTO } from './registry-dto'

/**
 *
 * @export
 * @interface BuildSnapshotRequestDTO
 */
export interface BuildSnapshotRequestDTO {
  /**
   *
   * @type {Array<string>}
   * @memberof BuildSnapshotRequestDTO
   */
  context?: Array<string>
  /**
   *
   * @type {string}
   * @memberof BuildSnapshotRequestDTO
   */
  dockerfile: string
  /**
   *
   * @type {string}
   * @memberof BuildSnapshotRequestDTO
   */
  organizationId: string
  /**
   *
   * @type {boolean}
   * @memberof BuildSnapshotRequestDTO
   */
  pushToInternalRegistry?: boolean
  /**
   *
   * @type {RegistryDTO}
   * @memberof BuildSnapshotRequestDTO
   */
  registry?: RegistryDTO
  /**
   * Snapshot ID and tag or the build\'s hash
   * @type {string}
   * @memberof BuildSnapshotRequestDTO
   */
  snapshot?: string
}

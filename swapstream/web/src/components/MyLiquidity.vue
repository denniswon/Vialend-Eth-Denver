<template>
  <div class="search_box">
    <div class="container">
      <div class="row">
        <div class="col">
          <div class="search_box_container d-flex flex-row align-items-center justify-content-start">
            <div class="search_form_container">
              <div class="my-liquidity-title">My Liquidity</div>
              <table class="table">
                <thead>
                  <th>Smart Vaults</th>
                  <th>Fee Tier</th>
                  <th>Current APR</th>
                  <th>Status</th>
                  <th>Current Value</th>
                  <th>&nbsp;</th>
                </thead>
                <tbody>
                  <tr>
                    <td>
                      <img src="images/usdc.png"
                           width="30"
                           height="30" />
                      <img src="images/weth.png"
                           width="30"
                           height="30" />
                      <span>USDC / WETH</span>
                    </td>
                    <td><input class="btn btn-default btn-sm"
                             type="button"
                             value="0.30%"></td>
                    <td>505.66%</td>
                    <td><button type="button"
                              class="btn btn-default btn-sm">In range</button></td>
                    <td>$33,500</td>
                    <td>Remove <button type="button"
                              class="btn btn-primary btn-sm"
                              disabled="true">Increase</button></td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MyLiquidity',
  methods: {

  }
}
</script>

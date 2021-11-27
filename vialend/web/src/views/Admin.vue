<template>
  <el-container>
    <el-header height="103px">
      <Header ref="headerComponents" />
    </el-header>
    <el-container>
      <el-aside background-color="#304156"
                width="200px">
        <el-menu default-active="10"
                 class="el-menu-vertical-demo"
                 @open="handleOpen"
                 @close="handleClose"
                 background-color="#304156"
                 text-color="#fff">
          <el-menu-item index="10"
                        @click="showContainer('template')">
            <i class="el-icon-menu"></i>
            <span slot="title">Template</span>
          </el-menu-item>
          <el-submenu index="20">
            <template slot="title">
              <i class="el-icon-s-help"></i>
              <span>Pairs</span>
            </template>
            <el-menu-item-group v-if="pairsList.size() > 0">
              <!-- :class="[menuItemClass,(selectedPairIndex === pair.index?' active':'')]" -->
              <el-menu-item v-for="pair in pairsList._getData()"
                            :key="pair.id"
                            :index="(21 + pair.index).toString()"
                            class="menu_text"
                            @click="changeSelectedPair(Number(pair.index))">
                <span>{{ pair.token0.symbol }} / {{ pair.token1.symbol }}</span>
              </el-menu-item>
            </el-menu-item-group>
          </el-submenu>
          <el-menu-item index="16"
                        @click="showContainer('setting')">
            <i class="el-icon-setting"></i>
            <span slot="title">Settings</span>
          </el-menu-item>
        </el-menu>
        <!-- <div class="list-group"
             v-if="pairsList.size() > 0">
          <a href="#"
             :class="[menuItemClass,(selectedPairIndex === pair.index?' active':'')]"
             v-for="pair in pairsList._getData()"
             :key="pair.id"
             @click="changeSelectedPair(pair.index)">
            <span>
              <img :src="pair.token0.iconLink"
                   width="20"
                   height="20" />&nbsp;
              <img :src="pair.token1.iconLink"
                   width="20"
                   height="20" />
            </span>&nbsp;&nbsp;
            <span><a href="#">{{ pair.token0.symbol }} / {{ pair.token1.symbol }}</a></span>
          </a>

        </div> -->
      </el-aside>
      <el-main>
        <!--Template Container-->
        <div id="template_container"
             :style="'display:' + (featureContainer === 'template' ? '' : 'none') "
             v-loading="templatePageLoading">
          <div style="float:right;padding:10px;">
            <el-button @click="showNewPairDialog"
                       plain>New Pair</el-button>&nbsp;&nbsp;
            <el-dropdown @command="handleCommand"
                         ref="messageDrop">
              <el-button>
                {{filterCommand}}<i class="el-icon-arrow-down el-icon--right"></i>
              </el-button>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item command="All">All</el-dropdown-item>
                <el-dropdown-item command="Active">Active</el-dropdown-item>
                <el-dropdown-item command="Pending">Pending</el-dropdown-item>
                <el-dropdown-item command="Closed">Closed</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </div>
          <div class="clear"></div>
          <el-divider></el-divider>
          <div style="padding:10px;"
               v-if="pairsList.size() > 0">
            <div class="block"
                 v-for="pair in pairsList._getData()"
                 :key="pair.id">
              <!-- <el-col :span="6"
                    :offset="2"
                    v-for="pair in pairsList._getData()"
                    :key="pair.id"> -->
              <el-card class="box-card"
                       v-loading="templatePairsLoading">
                <div class="text item">
                  <h3>{{ pair.token0.symbol }} / {{ pair.token1.symbol }}</h3>
                </div>
                <div class="text item">
                  Fee Tier<br><span style="color:#409eff"><input class="btn btn-default btn-sm"
                           type="button"
                           :value="Number(pair.feeTier / 10000) + '%'"></span>
                </div>
                <div class="text item">
                  Current APR<br><span style="color:#409eff">{{Number(pair.currentAPR).toFixed(2)}}%</span>
                </div>
                <div class="text item">
                  TVL<br><span style="color:#409eff">{{Number(pair.tvlTotal0).toFixed(2)}} / {{Number(pair.tvlTotal1).toFixed(2)}}<br>
                    <!-- ${{(Number(pair.tvlTotal0USD) + Number(pair.tvlTotal1USD)).toFixed(2)}} -->
                    </span>
                </div>
                <div class="text item">
                  Assets ratio<br><span style="color:#409eff">{{Number(pair.lendingRatio).toFixed(2)}}% Compound<br>
                    {{Number(pair.uniswapRatio).toFixed(2)}}% Uniswap V3</span>
                </div>
              </el-card>
              <!-- </el-col> -->
            </div>
          </div>
          <el-dialog title="New Pair"
                     :visible.sync="newPairDialogVisible"
                     :append-to-body="true"
                     width="600px"
                     :close-on-click-modal="false"
                     center>
            <el-form ref="newPairform"
                     :model="newPairform"
                     label-width="150px">

              <el-form-item label="Token0">
                <el-select v-model="newPairform.token0name"
                           placeholder="Please select token0">
                  <el-option label="WETH"
                             value="weth"></el-option>
                  <el-option label="USDC"
                             value="usdc"></el-option>
                  <el-option label="BTC"
                             value="btc"></el-option>
                  <el-option label="BNB"
                             value="bnb"></el-option>
                  <el-option label="SOL"
                             value="sol"></el-option>
                  <el-option label="USDT"
                             value="usdt"></el-option>
                  <el-option label="ADA"
                             value="ada"></el-option>
                  <el-option label="XRP"
                             value="xrp"></el-option>
                  <el-option label="DOT"
                             value="dot"></el-option>
                  <el-option label="UNI"
                             value="uni"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="Token1">
                <el-select v-model="newPairform.token1name"
                           placeholder="Please select token1">
                  <el-option label="WETH"
                             value="weth"></el-option>
                  <el-option label="USDC"
                             value="usdc"></el-option>
                  <el-option label="BTC"
                             value="btc"></el-option>
                  <el-option label="BNB"
                             value="bnb"></el-option>
                  <el-option label="SOL"
                             value="sol"></el-option>
                  <el-option label="USDT"
                             value="usdt"></el-option>
                  <el-option label="ADA"
                             value="ada"></el-option>
                  <el-option label="XRP"
                             value="xrp"></el-option>
                  <el-option label="DOT"
                             value="dot"></el-option>
                  <el-option label="UNI"
                             value="uni"></el-option>
                </el-select>
              </el-form-item>

              <el-form-item label="Fee Tier">
                <el-select v-model="newPairform.feetier"
                           placeholder="Please select fee tier">
                  <el-option label="0.05%"
                             value="0.05"></el-option>
                  <el-option label="0.3%"
                             value="0.3"></el-option>
                  <el-option label="1%"
                             value="1"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="UniswapV3 Portion">
                <el-select v-model="newPairform.unisV3Portion"
                           placeholder="Please select Uniswap V3 Portion">
                  <el-option label="0%"
                             value="0"></el-option>
                  <el-option label="20%"
                             value="20"></el-option>
                  <el-option label="50%"
                             value="50"></el-option>
                  <el-option label="80%"
                             value="80"></el-option>
                  <el-option label="100%"
                             value="100"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="Indicator">
                <el-select v-model="newPairform.rebalanceStrategy"
                           placeholder="Please select Indicator">
                  <el-option label="Bollinger"
                             value="Bollinger"></el-option>
                  <el-option label="Envelope"
                             value="Envelope"></el-option>
                  <el-option label="TMA range"
                             value="TMA range"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="Period">
                <el-select v-model="newPairform.period"
                           placeholder="Please select Period">
                  <el-option label="W1"
                             value="W1"></el-option>
                  <el-option label="D1"
                             value="D1"></el-option>
                  <el-option label="H4"
                             value="H4"></el-option>
                  <el-option label="H1"
                             value="H1"></el-option>
                  <el-option label="H15"
                             value="H15"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="Max Supply">
                <el-input v-model="newPairform.maxSupply"
                          style="width:200px;"></el-input>
              </el-form-item>

              <el-form-item label="Vault Address">
                <div style="float:left;">
                  <el-input v-model="newPairform.vauAddr"
                            style="width:360px;"></el-input>&nbsp;&nbsp;
                  <a :href="'https://goerli.etherscan.io/address/' + newPairform.vauAddr"
                     target="_blank"><img src="images/todetail.png"
                         width="20" /></a>
                </div>
              </el-form-item>

              <div style="display:none">
                <el-form-item label="NetworkId">
                  <el-input v-model="newPairform.networkid"></el-input>
                </el-form-item>
                <el-form-item label="Acc">
                  <el-input v-model="newPairform.acc"></el-input>
                </el-form-item>
                <el-form-item label="ProtocolFee">
                  <el-input v-model="newPairform.protocolfee"></el-input>
                </el-form-item>
                <el-form-item label="UniPortion">
                  <el-input v-model="newPairform.uniPortion"></el-input>
                </el-form-item>
                <el-form-item label="Team">
                  <el-input v-model="newPairform.team"></el-input>
                </el-form-item>
              </div>
              <el-form-item>
                <el-button type="primary"
                           @click="createNewPair"
                           :loading="deployLoading">Deploy</el-button>
                <el-button @click="newPairDialogVisible = false">Close</el-button>
              </el-form-item>
            </el-form>
          </el-dialog>
        </div>
        <!--Setting Container-->
        <div id="setting_container"
             :style="'display:' + (featureContainer === 'setting' ? '' : 'none')"
             v-loading="loadTickerLoading">
          <div class="token_exchange_form">
            <el-card class="box-card-currency">
              <div slot="header"
                   class="clearfix">
                <span>Currency Converter</span>
              </div>
              <div class="text item">
                <el-form :inline="true"
                         class="demo-form-inline">
                  <el-form-item>
                    <el-input v-model="currTokenVal"
                              placeholder="1"></el-input>
                  </el-form-item>
                  <el-form-item>
                    <el-select v-model="currTokenId"
                               placeholder="Please select"
                               @change="exchangeTokenChange">
                      <el-option v-for="item in tokensList"
                                 :key="item.nameid"
                                 :label="item.symbol"
                                 :value="item.symbol">
                      </el-option>
                    </el-select>
                    =
                  </el-form-item>
                  <el-form-item>
                    <el-input v-model="usdTokenVal"></el-input>
                  </el-form-item>
                  <el-form-item>
                    USD
                  </el-form-item>
                </el-form>
              </div>
            </el-card>
          </div>
        </div>
        <!--Pairs Container-->
        <div id="pairs_container"
             :style="'display:' + (featureContainer === 'pairs' ? '' : 'none') "
             v-loading="pairsInfoLoading">
          <!-- <div class="breadcrumb">
            <el-breadcrumb separator="/">
              <el-breadcrumb-item>
                <b>Current Pair:</b>
                <span style="color:#000000"
                      v-if="pairsList.size() > 0">
                  {{pairsList.get(selectedPairIndex).token0.symbol}} /
                  {{pairsList.get(selectedPairIndex).token1.symbol}}
                </span>
              </el-breadcrumb-item>
            </el-breadcrumb>
          </div> -->

          <el-tabs type="border-card">
            <el-tab-pane>
              <span slot="label"><i class="el-icon-date"></i> Rebalance</span>
              <div style="width:80%"
                   v-loading="loadingTokensInfoStatus">
                <table class="table table-striped">
                  <thead>
                    <tr>
                      <th>Token balance</th>
                      <th>Wallet</th>
                      <th>Vault</th>
                      <th>Uni-V3</th>
                      <th>Compound</th>
                      <!-- <th>fees</th> -->
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <th>{{token0Symbol}}</th>
                      <td>{{Number(token0BalanceInWallet).toFixed(2)}}</td>
                      <!-- + ' / $' + token0BalanceUSDInWallet.toFixed(2) -->
                      <td>{{Number(token0BalanceInVault).toFixed(2)}}</td>
                      <!-- + ' / $' + token0BalanceUSDInVault.toFixed(2) -->
                      <td>{{Number(token0BalanceInPool).toFixed(2)}}</td>
                      <!-- + ' / $' + token0BalanceUSDInPool.toFixed(2) -->
                      <td>{{Number(token0BalanceInLending).toFixed(2)}}</td>
                      <!-- + ' / $' + token0BalanceUSDInLending.toFixed(2) -->
                      <!-- <td>{{accruedProtocolFees0}}</td> -->
                    </tr>
                    <tr>
                      <th>{{token1Symbol}}</th>
                      <td>{{Number(token1BalanceInWallet).toFixed(2)}}</td>
                      <!-- + ' / $' + token1BalanceUSDInWallet.toFixed(2) -->
                      <td>{{Number(token1BalanceInVault).toFixed(2)}}</td>
                      <!-- + ' / $' + token1BalanceUSDInVault.toFixed(2) -->
                      <td>{{Number(token1BalanceInPool).toFixed(2)}}</td>
                      <!-- + ' / $' + token1BalanceUSDInPool.toFixed(2) -->
                      <td>{{Number(token1BalanceInLending).toFixed(2)}}</td>
                      <!-- + ' / $' + token1BalanceUSDInLending.toFixed(2) -->
                      <!-- <td>{{accruedProtocolFees1}}</td> -->
                    </tr>
                    <!-- <tr>
                      <th>USD</th>
                      <td>${{(Number(token0BalanceUSDInWallet) + Number(token1BalanceUSDInWallet)).toFixed(2)}}</td>
                      <td>${{(Number(token0BalanceUSDInVault) + Number(token1BalanceUSDInVault)).toFixed(2)}}</td>
                      <td>${{(Number(token0BalanceUSDInPool) + Number(token1BalanceUSDInPool)).toFixed(2)}}</td>
                      <td>${{(Number(token0BalanceUSDInLending) + Number(token1BalanceUSDInLending)).toFixed(2)}}</td>
                       <td>{{accruedProtocolFees1}}</td>
                    </tr> -->
                  </tbody>
                </table>
              </div>
              <hr>
              Min Tick:{{tickLower}}<br>
              Max Tick:{{tickUpper}}<br>
              Current Tick:{{currentTick}}<br>
              <span style="color:#ff0000;margin-right:10px;">Current Price:{{currentPrice}}</span>
              <el-button size="mini"
                         :style="{display:rangeStatusDisplay}">
                <span :class="[dotStyle,rangeStatusStyle]"></span>
                <span class="status">{{rangeStatus}}</span>
              </el-button><br>
              <div class="block"
                   style="width:80%;margin-top:20px;margin-left:50px">
                <!-- <el-slider v-model="tickRange"
                     range
                     :marks="getMarks">
          </el-slider> -->
                <input type="text"
                       class="js-range-slider"
                       name="my_range"
                       value="" />
              </div>
              <div class="clear"></div>
              <div class="range_title">Set Price Range</div>
              <el-form :inline="true"
                       :model="rangeForm"
                       class="demo-form-inline">
                <el-form-item label="Min Price">
                  <el-input-number v-model="rangeForm.minPrice"
                                   size="medium"
                                   @change="minPriceChange"
                                   :precision="precision"
                                   :step="step"></el-input-number><br>
                  {{token1Symbol}} per {{token0Symbol}}
                </el-form-item>
                <el-form-item label="Max Price">
                  <el-input-number v-model="rangeForm.maxPrice"
                                   size="medium"
                                   @change="maxPriceChange"
                                   :precision="precision"
                                   :step="step"></el-input-number><br>
                  {{token1Symbol}} per {{token0Symbol}}
                </el-form-item>
                <br>
                <el-form-item>
                  <el-button type="primary"
                             :loading="rebalanceLoading"
                             @click="doRebalance">Rebalance</el-button>
                  <!-- <el-button type="primary"
                       @click="getY">getY</el-button> -->
                  &nbsp;&nbsp;
                  <a :href="goToEtherscan"
                     target="_blank"
                     :class="['btn','btn-primary',doRebalanceEtherscanDisable?'disabled':'']"
                     role="button">View on etherscan</a>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary"
                             :loading="removePositionsLoading"
                             @click="doRemovePositions">RemovePositions</el-button>
                  <!-- <el-button type="primary"
                       @click="getY">getY</el-button> -->
                  &nbsp;&nbsp;
                  <a :href="goToEtherscan"
                     target="_blank"
                     :class="['btn','btn-primary',removePositionsEtherscanDisable?'disabled':'']"
                     role="button">View on etherscan</a>
                </el-form-item>
              </el-form>
              {{errorRebalance}}
            </el-tab-pane>
            <el-tab-pane>
              <span slot="label"><i class="el-icon-date"></i> Security</span>
              <el-form :inline="true"
                       class="demo-form-inline">
                <table class="table table-bordered table-hover">
                  <tr>
                    <td>
                      <el-form-item>Pool Address:</el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        {{poolAddress}}&nbsp;&nbsp;<a :href="'https://goerli.etherscan.io/address/' + poolAddress"
                           target="_blank"><img src="images/todetail.png"
                               width="20" /></a>
                      </el-form-item>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <el-form-item>Vault Address:</el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        {{vaultAddress}}&nbsp;&nbsp;<a :href="'https://goerli.etherscan.io/address/' + vaultAddress"
                           target="_blank"><img src="images/todetail.png"
                               width="20" /></a>
                      </el-form-item>
                    </td>
                  </tr>
                  <tr>
                    <td colspan="2">
                      <el-form-item>
                        <el-button type="warning"
                                   :loading="emergencyLoading"
                                   @click="EmergencyBurn">EmergencyBurn</el-button>
                        &nbsp;&nbsp;
                        <a :href="goToEtherscan"
                           target="_blank"
                           :class="['btn','btn-primary',emergencyEtherscanDisable?'disabled':'']"
                           role="button">View on etherscan</a>
                      </el-form-item>
                    </td>
                  </tr>
                </table>
              </el-form>
            </el-tab-pane>
            <el-tab-pane>
              <span slot="label"><i class="el-icon-date"></i> Settings</span>
              <!-- <div class="token_exchange_form">
              <div style="margin-bottom:20px;">Currency Converter</div>
              <el-form :inline="true"
                       class="demo-form-inline">
                <el-form-item>
                  <el-input v-model="currTokenVal"
                            placeholder="1"></el-input>
                </el-form-item>
                <el-form-item>
                  <el-select v-model="currTokenId"
                             placeholder="请选择"
                             @change="exchangeTokenChange">
                    <el-option v-for="item in tokensList"
                               :key="item.nameid"
                               :label="item.symbol"
                               :value="item.symbol">
                    </el-option>
                  </el-select>
                </el-form-item>
                =
                <el-form-item>
                  <el-input v-model="usdTokenVal"></el-input>
                </el-form-item>USD
              </el-form>
            </div>
            <hr> -->
              <el-form :inline="true"
                       class="demo-form-inline">
                <table class="table table-bordered table-hover">
                  <tr>
                    <td>
                      <el-form-item>UniPortionRatio:</el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-input v-model="uniPortionRatio"
                                  style="width:500px;"></el-input>
                      </el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-button type="primary"
                                   :loading="uniPortionRatioLoading"
                                   @click="SetUniPortionRatio">SetUniPortionRatio</el-button>
                        &nbsp;&nbsp;
                        <a :href="goToEtherscan"
                           target="_blank"
                           :class="['btn','btn-primary',uniPortionRatioEtherscanDisable?'disabled':'']"
                           role="button">View on etherscan</a>
                      </el-form-item>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <el-form-item>MaxTotalSupply:</el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-input v-model="maxTotalSupply"
                                  style="width:500px;"></el-input>
                      </el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-button type="primary"
                                   :loading="maxTotalSupplyLoading"
                                   @click="SetMaxTotalSupply">SetMaxTotalSupply</el-button>
                        &nbsp;&nbsp;
                        <a :href="goToEtherscan"
                           target="_blank"
                           :class="['btn','btn-primary',maxTotalSupplyEtherscanDisable?'disabled':'']"
                           role="button">View on etherscan</a>
                      </el-form-item>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <el-form-item>ProtocolFee:</el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-input v-model="protocolFee"
                                  style="width:500px;"></el-input>
                      </el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-button type="primary"
                                   :loading="protocolFeeLoading"
                                   @click="SetProtocolFee">SetProtocolFee</el-button>
                        &nbsp;&nbsp;
                        <a :href="goToEtherscan"
                           target="_blank"
                           :class="['btn','btn-primary',protocolFeeEtherscanDisable?'disabled':'']"
                           role="button">View on etherscan</a>
                      </el-form-item>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <el-form-item>Governance:</el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-input v-model="governanceAddress"
                                  style="width:500px;"></el-input>&nbsp;&nbsp;<a :href="'https://goerli.etherscan.io/address/' + governanceAddress"
                           target="_blank"><img src="images/todetail.png"
                               width="20" /></a>
                      </el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-button type="primary"
                                   :loading="governanceLoading"
                                   @click="SetGovernance">SetGovernance</el-button>
                        <el-button type="primary"
                                   :loading="acceptGovernanceLoading"
                                   @click="acceptGovernance">AccecptGovernance</el-button>
                        &nbsp;&nbsp;
                        <a :href="goToEtherscan"
                           target="_blank"
                           :class="['btn','btn-primary',governanceEtherscanDisable?'disabled':'']"
                           role="button">View on etherscan</a>
                      </el-form-item>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <el-form-item>Team:</el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-input v-model="teamAddress"
                                  style="width:500px;"></el-input>&nbsp;&nbsp;<a :href="'https://goerli.etherscan.io/address/' + teamAddress"
                           target="_blank"><img src="images/todetail.png"
                               width="20" /></a>
                      </el-form-item>
                    </td>
                    <td>
                      <el-form-item>
                        <el-button type="primary"
                                   :loading="teamLoading"
                                   @click="SetTeam">SetTeam</el-button>
                        &nbsp;&nbsp;
                        <a :href="goToEtherscan"
                           target="_blank"
                           :class="['btn','btn-primary',teamEtherscanDisable?'disabled':'']"
                           role="button">View on etherscan</a>
                      </el-form-item>
                    </td>
                  </tr>
                </table>
              </el-form>
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import Web3 from 'web3'
import Header from '@/components/Header.vue'
import contractABI from '../ABI/ViaLendFeeMakerABI.json'
import uniswapV3PoolABI from '../ABI/UniswapV3PoolABI.json'
import ViaLendTokenABI from '../ABI/tokenABI.json'
import VaultBridgeABI from '../ABI/VaultBridge.json'
import Token from '../model/Token'
import Pairs from '../model/Pairs'
import ArrayList from '../common/ArrayList'
import axios from 'axios'

if (typeof web3 !== 'undefined') {
  web3 = new Web3(web3.currentProvider)
  console.log('web3 provider:web3.currentProvider')
} else {
  // set the provider you want from Web3.providers
  web3 = new Web3(new Web3.providers.HttpProvider('https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6'))
  console.log('web3 provider:goerli')
}

export default {
  components: { Header },
  data () {
    return {
      featureContainer: 'template',
      exchangeTimer: '',
      pairsList: new ArrayList(),
      selectedPairIndex: 0,
      isConnected: this.$store.state.isConnected,
      keeperContract: null,
      poolAddress: '',
      vaultAddress: '',
      tickLower: 0,
      tickUpper: 0,
      currentTick: 0,
      currentPrice: 0,
      tickRange: [20, 80],
      currentPercent: 0,
      precision: 18,
      step: 0.000001,
      marks: {
      },
      rangeForm: {
        minPrice: 0.0,
        maxPrice: 0.0
      },
      newPairform: {
        networkid: 3,
        acc: 0,
        protocolfee: 10,
        uniPortion: 30,
        team: '0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440',
        token0name: '',
        token1name: '',
        feetier: null,
        unisV3Portion: '',
        rebalanceStrategy: '',
        maxSupply: '',
        vauAddr: '',
        period: ''
      },
      errorRebalance: '',
      // Control's load state
      rebalanceLoading: false,
      removePositionsLoading: false,
      emergencyLoading: false,
      protocolFeeLoading: false,
      teamLoading: false,
      uniPortionRatioLoading: false,
      maxTotalSupplyLoading: false,
      governanceLoading: false,
      acceptGovernanceLoading: false,
      deployLoading: false,
      loadTickerLoading: false,
      templatePageLoading: false,
      templatePairsLoading: false,
      pairsInfoLoading: false,
      // Control's enable state
      doRebalanceEtherscanDisable: true,
      removePositionsEtherscanDisable: true,
      emergencyEtherscanDisable: true,
      protocolFeeEtherscanDisable: true,
      uniPortionRatioEtherscanDisable: true,
      maxTotalSupplyEtherscanDisable: true,
      governanceEtherscanDisable: true,
      teamEtherscanDisable: true,
      loadingTokensInfoStatus: false,
      token0Contract: null,
      token1Contract: null,
      token0LendingContract: null,
      token1LendingContract: null,
      token0Address: '',
      token0Name: '',
      token0Symbol: '',
      token0Decimal: 0,
      token0Balance: 0,
      token1Address: '',
      token1Name: '',
      token1Symbol: '',
      token1Decimal: 0,
      token0BalanceInWallet: 0.00,
      token1BalanceInWallet: 0.00,
      token0BalanceUSDInWallet: 0.00,
      token1BalanceUSDInWallet: 0.00,
      token0BalanceInVault: 0.00,
      token1BalanceInVault: 0.00,
      token0BalanceUSDInVault: 0.00,
      token1BalanceUSDInVault: 0.00,
      token0BalanceInPool: 0,
      token1BalanceInPool: 0,
      token0BalanceUSDInPool: 0,
      token1BalanceUSDInPool: 0,
      token0BalanceInLending: 0,
      token1BalanceInLending: 0,
      token0BalanceUSDInLending: 0,
      token1BalanceUSDInLending: 0,
      menuItemClass: 'list-group-item',
      dotStyle: 'dot',
      rangeStatusStyle: '',
      rangeStatusDisplay: 'none',
      rangeStatus: '',
      tokensList: null,
      currTokenId: 'BTC',
      currTokenVal: 1,
      usdTokenVal: 0,
      priceUSD: 0,
      maxTotalSupply: 0,
      uniPortionRatio: 0,
      governanceAddress: '',
      teamAddress: '',
      uniPortion: 0,
      protocolFee: 0,
      viewOnEtherscanDisable: true,
      transactionHash: '',
      goToEtherscan: '',
      accruedProtocolFees0: 0,
      accruedProtocolFees1: 0,
      priceRangeSlider: null,
      newPairDialogVisible: false,
      filterCommand: 'All',
      bridgeAddress: '0x033F3C5eAd18496BA462783fe9396CFE751a2342',
      pairsLoadComplete: false
    }
  },
  created: async function () {
    if (ethereum.selectedAddress !== null && ethereum.selectedAddress !== undefined) {
      if (!this.$store.state.isAdmin) {
        if (localStorage.getItem('isAdmin') == null || localStorage.getItem('isAdmin') === undefined) {
          console.log('localStorage isAdmin is null')
          self.location.replace('/dashboard')
        } else if (!localStorage.getItem('isAdmin')) {
          console.log('localStorage isAdmin is false')
          self.location.replace('/dashboard')
        }
      }
    } else if (localStorage.getItem('isAdmin') == null || localStorage.getItem('isAdmin') === undefined) {
      console.log('localStorage isAdmin is null')
      self.location.replace('/dashboard')
      // return
    } else if (!localStorage.getItem('isAdmin')) {
      console.log('localStorage isAdmin is false')
      self.location.replace('/dashboard')
    }
    this.templatePageLoading = true

    console.log('validNetwork=', this.$store.state.validNetwork, ';pairsList size=', this.pairsList.size())
    console.log('parent pairsList size=', this.$parent.pairsList.size())
    if (this.$store.state.validNetwork) {
      if (this.$parent.pairsList.size() === 0) {
        this.$parent.loadPairsInfo()
      } else {
        this.pairsList = this.$parent.pairsList
        this.loadTemplateData()
      }
    }
    // if (!this.$store.state.validNetwork) {
    //   for (var i = 0; i < this.pairsList.size(); i++) {
    //     this.pairsList.get(i).Empty()
    //     // console.log('pair feeTier=', this.pairsList.get(i).feeTier)
    //   }
    //   console.log('Chain is unavailable,clear pairsList:', this.pairsList.size())
    //   // to be used
    //   this.myValueToken0Locked = 0.00
    //   this.myValueToken1Locked = 0.00
    //   this.myValueToken0USDLocked = 0
    //   this.myValueToken1USDLocked = 0
    //   this.token0BalanceInWallet = 0.00
    //   this.token1BalanceInWallet = 0.00
    //   this.token0BalanceUSDInWallet = 0.00
    //   this.token1BalanceUSDInWallet = 0.00
    //   this.token0BalanceInVault = 0.00
    //   this.token1BalanceInVault = 0.00
    //   this.token0BalanceUSDInVault = 0.00
    //   this.token1BalanceUSDInVault = 0.00
    //   this.token0BalanceInPool = 0.00
    //   this.token1BalanceInPool = 0.00
    //   this.token0BalanceUSDInPool = 0.00
    //   this.token1BalanceUSDInPool = 0.00
    //   this.token0BalanceInLending = 0.00
    //   this.token1BalanceInLending = 0.00
    //   this.token0BalanceUSDInLending = 0.00
    //   this.token1BalanceUSDInLending = 0.00
    // } else {
    //   console.log('pairsList size=', this.$parent.pairsList.size())
    //   // await this.loadPairsInfo()
    //   // this.loadTemplateData()
    //   // this.initData()
    // }

    // this.getTokensBalanceInWallet()
    // //
    // var result = await keeperContract.methods.getPositionAmounts(BigInt(pair.cLow), BigInt(pair.cHigh)).call()
    // if (result !== undefined && result !== null) {
    //   this.token0BalanceInPool = result.amount0 / Number(Math.pow(10, this.token0Decimal))
    //   this.token0BalanceUSDInPool = Number(this.token0BalanceInPool) * Number(this.$store.state.token0RateOfUSD)
    //   this.token1BalanceInPool = result.amount1 / Number(Math.pow(10, this.token1Decimal))
    //   this.token1BalanceUSDInPool = Number(this.token1BalanceInPool) * Number(this.$store.state.token1RateOfUSD)
    // }
    this.templatePageLoading = false
  },
  mounted () {
    this.$refs.headerComponents.containerClass = 'container-fluid'
    this.exchangeTimer = setInterval(this.exchangeTokensIntoUSD, 1000)
  },
  computed: {
    newMinPrice () {
      return this.rangeForm.minPrice
    },
    newMaxPrice () {
      return this.rangeForm.maxPrice
    },
    getMarks () {
      var marks = {}
      if (this.currentPercent !== 0) {
        marks[this.currentPercent] = this.currentTick
      }
      return marks
    },
    pairsListSize () {
      return this.$parent.pairsList.size()
    },
    parentPairsListComplete () {
      return this.$parent.pairsLoadComplete
    }
  },
  beforeDestroy () {
    clearInterval(this.exchangeTimer)
  },
  watch: {
    // pairsListSize (size) {
    //   if (size > 0) {
    //     console.log('currentPairsList size=', size, ';validNetwork=', this.$store.state.validNetwork)
    //     if (this.$parent.pairsLoadComplete) {
    //       this.pairsList = this.$parent.pairsList
    //       this.pairsLoadComplete = true
    //     }
    //   }
    // },
    // pairsLoadComplete: function (newVal, oldVal) {
    //   console.log('pairs base info load complete：', newVal, ';old status:', oldVal)
    //   if (this.$store.state.isConnected && this.$store.state.validNetwork && !this.myPairsListLoading) {
    //     this.loadTemplateData()
    //   }
    // },
    parentPairsListComplete (newVal) {
      console.log('pairs base info load complete：', newVal)
      if (newVal) {
        if (this.$store.state.isConnected && this.$store.state.validNetwork && !this.myPairsListLoading) {
          console.log('begin loadTemplateData')
          this.pairsList = this.$parent.pairsList
          this.loadTemplateData()
        }
      }
    },
    newMinPrice (price) {
      console.log('new min price=', price)
      if (!isNaN(price)) {
        console.log('tickLower->price:', price)
        this.tickLower = this.priceToTick(price, this.token0Decimal, this.token1Decimal)
      } else { this.tickLower = 0 }
      this.drawTickRangeChart()
    },
    newMaxPrice (price) {
      console.log('new max price=', price)
      if (!isNaN(price)) {
        console.log('tickUpper->price:', price)
        this.tickUpper = this.priceToTick(price, this.token0Decimal, this.token1Decimal)
      } else { this.tickUpper = 0 }
      this.drawTickRangeChart()
    },
    '$store.state.isConnected': function (newVal, oldVal) {
      console.log('connection status,new:', newVal, 'old:', oldVal)
      this.isConnected = this.$store.state.isConnected
      // if (this.isConnected) {
      //   console.log('wallet connected')
      //   this.getTokensBalanceInWallet()
      // } else {
      //   this.token0BalanceInWallet = 0.00
      //   this.token1BalanceInWallet = 0.00
      //   this.token0BalanceUSDInWallet = 0.00
      //   this.token1BalanceUSDInWallet = 0.00
      // }
    },
    '$store.state.currentAccount': function (newVal, oldVal) {
      // if (!this.$store.state.isAdmin) {
      //   self.location.replace('/dashboard')
      //   return
      // }
      console.log('currentAccount:', newVal, ';previousAccount:', oldVal)
      if (newVal !== '' && this.$store.state.validNetwork) {
        console.log('Account changed,pairlist size:', this.pairsList.size())
        if (this.pairsList.size() === 0 && !this.$parent.pairsInfoLoading) {
          this.$parent.loadPairsInfo()
        } else {
          this.pairsListComplete = false
          if (this.featureContainer === 'template') {
            this.loadTemplateData()
          } else if (this.featureContainer === 'pairs') {
            this.loadTokensBalance()
          }
        }
      }
    },
    '$store.state.chainId': function () {
      console.log('chain changed')
      // this.initData()
    },
    currTokenVal (val) {
      if (val !== '') {
        if (!isNaN(val)) {
          console.log('currTokenVal=', val)
          this.usdTokenVal = val * this.priceUSD
        }
      }
    },
    currentTick (val) {
      this.drawTickRangeChart()
    }
  },
  methods: {
    contractInstance (abi, address) {
      return new web3.eth.Contract(abi, address)
    },
    showContainer (id) {
      if (id === 'template') {
        this.featureContainer = 'template'
        this.loadTemplateData()
      } else if (id === 'pairs') {
        this.featureContainer = 'pairs'
      } else if (id === 'setting') {
        this.featureContainer = 'setting'
        this.loadSettingData()
      }
    },
    showNewPairDialog () {
      this.newPairDialogVisible = true
    },
    async createNewPair () {
      var _this = this
      this.deployLoading = true
      // var url = 'http://localhost:8081/deploy?networkid=' + this.newPairform.networkid + '&acc=' + this.newPairform.acc + '&protocolfee=' + this.newPairform.protocolfee + '&uniPortion=' + this.newPairform.uniPortion + '&team=' + this.newPairform.team
      var url = 'http://118.126.117.7:8081/deploy?networkid=' + this.newPairform.networkid + '&acc=' + this.newPairform.acc + '&protocolfee=' + this.newPairform.protocolfee + '&uniPortion=' + this.newPairform.uniPortion + '&team=' + this.newPairform.team
      console.log('url=', url)
      axios.get(url).then(function (response) {
        if (response !== undefined && response !== null) {
          console.log('response=', response.data)
          if (response.data !== undefined && response.data !== null) {
            //   this.$store.state.tokenExchangeTable[i].price_usd = response.data[0].price_usd
            _this.newPairform.vauAddr = response.data['vault address']
          }
        }
        _this.deployLoading = false
      })
    },
    handleCommand (command) {
      // this.$message('click on item ' + command)
      this.filterCommand = command
      this.$refs.messageDrop.hide()
    },
    handleOpen (key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose (key, keyPath) {
      console.log(key, keyPath)
    },
    async loadTemplateData () {
      this.templatePairsLoading = true
      for (var i = 0; i < this.pairsList.size(); i++) {
        await this.getPairPublicInfo(this.pairsList.get(i))
      }
      this.templatePairsLoading = false
    },
    async loadSettingData () {
      this.loadTickerLoading = true
      var response = await axios.get('https://api.coinlore.net/api/tickers/')
      if (response !== undefined && response !== null) {
        console.log('response=', response.data.data)
        if (response.data.data.length > 0) {
          this.tokensList = response.data.data
          this.priceUSD = this.tokensList[0].price_usd
          this.usdTokenVal = this.priceUSD
          console.log('this.tokensList=', this.tokensList)
        }
      }
      this.loadTickerLoading = false
    },
    getTokenRateOfUSD (symbol) {
      // console.log('symbol=', symbol)
      if (this.$store.getters.getPriceUSDBySymbol(symbol) !== undefined) {
        // console.log('eth exchange table of ETH:', this.$store.getters.getPriceUSDBySymbol(symbol).price_usd)
        return this.$store.getters.getPriceUSDBySymbol(symbol).price_usd
      } else {
        return -1
      }
    },
    exchangeTokensIntoUSD () {
      if (this.$store.state.tokenExchangeRateLoaded && this.pairsList.get(this.selectedPairIndex) !== undefined) {
        var token0RateOfUSD, token1RateOfUSD
        // console.log('this.pairsList.get(th)=', this.pairsList.get(this.selectedPairIndex))
        token0RateOfUSD = this.getTokenRateOfUSD(this.pairsList.get(this.selectedPairIndex).token0.symbol)
        token1RateOfUSD = this.getTokenRateOfUSD(this.pairsList.get(this.selectedPairIndex).token1.symbol)
        // console.log('token0 RateOfUSD=', token0RateOfUSD)
        // console.log('token1 RateOfUSD=', token1RateOfUSD)
        this.token0BalanceUSDInWallet = Number(this.token0BalanceInWallet) * Number(token0RateOfUSD)
        this.token1BalanceUSDInWallet = Number(this.token1BalanceInWallet) * Number(token1RateOfUSD)
        this.token0BalanceUSDInVault = Number(this.token0BalanceInVault) * Number(token0RateOfUSD)
        this.token1BalanceUSDInVault = Number(this.token1BalanceInVault) * Number(token1RateOfUSD)
        this.token0BalanceUSDInPool = Number(this.token0BalanceInPool) * Number(token0RateOfUSD)
        this.token1BalanceUSDInPool = Number(this.token1BalanceInPool) * Number(token1RateOfUSD)
        this.token0BalanceUSDInLending = Number(this.token0BalanceInLending) * Number(token0RateOfUSD)
        this.token1BalanceUSDInLending = Number(this.token1BalanceInLending) * Number(token1RateOfUSD)

        for (var i = 0; i < this.pairsList.size(); i++) {
          token0RateOfUSD = this.getTokenRateOfUSD(this.pairsList.get(i).token0.symbol)
          token1RateOfUSD = this.getTokenRateOfUSD(this.pairsList.get(i).token1.symbol)
          // console.log('token0 RateOfUSD=', token0RateOfUSD)
          // console.log('token1 RateOfUSD=', token1RateOfUSD)
          // if (this.pairsList.get(i).tvlTotal0 >= 0) this.pairsList.get(i).tvlTotal0USD = Number(this.pairsList.get(i).tvlTotal0) * Number(token0RateOfUSD)
          // if (this.pairsList.get(i).tvlTotal1 >= 0) this.pairsList.get(i).tvlTotal1USD = Number(this.pairsList.get(i).tvlTotal1) * Number(token1RateOfUSD)
          this.pairsList.get(i).tvl = this.pairsList.get(i).tvlTotal0USD + this.pairsList.get(i).tvlTotal1USD
          if (this.pairsList.get(i).myValueToken0Locked >= 0) {
            this.pairsList.get(i).myValueToken0USDLocked = Number(this.pairsList.get(i).myValueToken0Locked) * Number(token0RateOfUSD)
          }
          if (this.pairsList.get(i).myValueToken1Locked >= 0) {
            this.pairsList.get(i).myValueToken1USDLocked = Number(this.pairsList.get(i).myValueToken1Locked) * Number(token1RateOfUSD)
          }
        }
      }
      // console.log('tokenExchangeRateLoaded:', this.$store.state.tokenExchangeRateLoaded)
    },
    changeSelectedPair (index) {
      this.featureContainer = 'pairs'
      console.log('current pair index=', index)
      this.selectedPairIndex = index
      // // this.loadingTokensInfoStatus = true
      if (this.pairsList.size() > index) {
        this.vaultAddress = this.pairsList.get(index).vaultAddress
        this.keeperContract = this.contractInstance(contractABI, this.pairsList.get(index).vaultAddress)
        this.poolAddress = this.pairsList.get(index).poolAddress
        this.token0Address = this.pairsList.get(index).token0.tokenAddress
        this.token1Address = this.pairsList.get(index).token1.tokenAddress
        this.token0Name = this.pairsList.get(index).token0.name
        this.token0Symbol = this.pairsList.get(index).token0.symbol
        this.token0Decimal = this.pairsList.get(index).token0.decimals
        this.token1Name = this.pairsList.get(index).token1.name
        this.token1Symbol = this.pairsList.get(index).token1.symbol
        this.token1Decimal = this.pairsList.get(index).token1.decimals
        if (Number(this.pairsList.get(index).cLow) > 0 && Number(this.pairsList.get(index).cHigh) > 0) {
          this.tickLower = this.pairsList.get(index).cLow
          this.tickUpper = this.pairsList.get(index).cHigh
          console.log('changeSelectedPair=>tickLower', this.tickLower)
          console.log('changeSelectedPair=>tickUpper', this.tickUpper)
        } else {
          this.tickLower = 0
          this.tickUpper = 0
        }
        this.token0LendingAddress = this.pairsList.get(index).token0.token0LendingAddress
        this.token1LendingAddress = this.pairsList.get(index).token1.token1LendingAddress
        this.loadPairsData()
      }
    },
    async loadPairsData () {
      this.pairsInfoLoading = true
      this.loadTokensBalance()
      this.getPairsSettingData()
      this.pairsInfoLoading = false
    },
    async loadTokensBalance () {
      var cLow, cHigh
      this.loadingTokensInfoStatus = true
      this.token0LendingContract = await this.contractInstance(ViaLendTokenABI, this.token0LendingAddress)
      this.token1LendingContract = await this.contractInstance(ViaLendTokenABI, this.token1LendingAddress)
      console.log('token0LendingAddress=', this.token0LendingAddress, 'token1LendingAddress=', this.token1LendingAddress)
      // ----------------------Get Slot0 Information-------------------------
      var poolAddress = await this.keeperContract.methods.poolAddress().call()
      var keeperUniswapV3Contract = this.contractInstance(uniswapV3PoolABI, poolAddress)
      var slot0 = await keeperUniswapV3Contract.methods.slot0().call()
      if (slot0 !== null && slot0 !== undefined) {
        this.currentTick = slot0['tick']
        this.currentPrice = this.tickToPrice(this.currentTick, this.token0Decimal, this.token1Decimal)
        if (this.tickLower === 0 || this.tickUpper === 0) {
          this.tickLower = parseInt(this.currentTick) - 500
          this.tickUpper = parseInt(this.currentTick) + 500
          console.log('slot0_tickLower0=', this.tickLower, 'slot0_tickUpper0=', this.tickUpper)
        } else {
          console.log('slot0_tickLower1=', this.tickLower, ';slot0_tickUpper1=', this.tickUpper)
        }
      }
      // ----------------------Tokens balance ------------------------------
      // Token0 balance in wallet and vault
      this.getTokensBalanceInWallet()
      // Token0 and token1 balance in pool
      cLow = await this.keeperContract.methods.cLow().call()
      cHigh = await this.keeperContract.methods.cHigh().call()

      console.log('cLow=', cLow, 'cHigh=', cHigh)
      var result = await this.keeperContract.methods.getPositionAmounts(Number(cLow), Number(cHigh)).call()
      console.log('result=', JSON.stringify(result))
      if (result !== undefined && result !== null) {
        this.token0BalanceInPool = result.amount0 / Number(Math.pow(10, this.token0Decimal))
        this.token0BalanceUSDInPool = Number(this.token0BalanceInPool) * Number(this.$store.state.token0RateOfUSD)
        this.token1BalanceInPool = result.amount1 / Number(Math.pow(10, this.token1Decimal))
        this.token1BalanceUSDInPool = Number(this.token1BalanceInPool) * Number(this.$store.state.token1RateOfUSD)
        console.log('token0BalanceInPool=', this.token0BalanceInPool, ';token1BalanceInPool=', this.token1BalanceInPool)
      }
      // Token0 and token1 balance in Compound
      var exchangeRate0 = await this.token0LendingContract.methods.exchangeRateStored().call()
      var exchangeRate1 = await this.token1LendingContract.methods.exchangeRateStored().call()
      var CAmount0 = await this.token0LendingContract.methods.balanceOf(this.vaultAddress).call()
      var CAmount1 = await this.token1LendingContract.methods.balanceOf(this.vaultAddress).call()
      var underlying0 = CAmount0 * exchangeRate0 / Math.pow(10, 18)
      var underlying1 = CAmount1 * exchangeRate1 / Math.pow(10, 18)
      console.log('vaultAddress=', this.vaultAddress)
      console.log('exchangeRate0=', exchangeRate0, 'exchangeRate1=', exchangeRate1, 'CAmount0=', CAmount0, 'CAmount1=', CAmount1, 'underlying0=', underlying0, 'underlying1=', underlying1)
      if (!isNaN(underlying0) && !isNaN(underlying1)) {
        this.token0BalanceInLending = underlying0 / Number(Math.pow(10, this.token0Decimal))
        this.token0BalanceUSDInLending = Number(this.token0BalanceInLending) * Number(this.$store.state.token0RateOfUSD)
        this.token1BalanceInLending = underlying1 / Number(Math.pow(10, this.token1Decimal))
        console.log('underlying1=', underlying1, 'token1BalanceInLending=', this.token1BalanceInLending)
        this.token1BalanceUSDInLending = Number(this.token1BalanceInLending) * Number(this.$store.state.token1RateOfUSD)
      }
      // Load price range
      if (Number(this.tickLower) === 0 && Number(this.tickUpper) === 0) {
        var clow = await this.keeperContract.methods.cLow().call()
        var chigh = await this.keeperContract.methods.cHigh().call()
        if (Number(clow) > 0 && Number(chigh) > 0) {
          this.tickLower = clow
          this.tickUpper = chigh
          console.log('loadTokensBalance=>tickLower=', this.tickLower, 'loadTokensBalance=>tickUpper=', this.tickUpper)
        }
      }
      console.log('tickLower_price=', this.tickLower, ';this.token0Decimal=', this.token0Decimal, ';this.token1Decimal=', this.token1Decimal)
      console.log('tickUpper_price=', this.tickUpper)
      this.rangeForm.minPrice = this.tickToPrice(this.tickLower, this.token0Decimal, this.token1Decimal)
      this.rangeForm.maxPrice = this.tickToPrice(this.tickUpper, this.token0Decimal, this.token1Decimal)
      console.log('rangeForm.minPrice=', this.rangeForm.minPrice, 'rangeForm.maxPrice=', this.rangeForm.maxPrice)
      var _this = this
      $('.js-range-slider').ionRangeSlider({
        skin: 'big',
        type: 'double',
        grid: true,
        prefix: '$',
        step: this.step,
        onChange: function (data) {
          console.log('from=', data.from)
          _this.rangeForm.minPrice = data.from
          _this.rangeForm.maxPrice = data.to
        }
      })
      this.priceRangeSlider = $('.js-range-slider').data('ionRangeSlider')
      this.priceRangeSlider.update({
        min: (parseFloat(this.rangeForm.minPrice) - 230) < 0 ? 0 : (parseFloat(this.rangeForm.minPrice) - 230).toFixed(1),
        max: (parseFloat(this.rangeForm.maxPrice) + 230).toFixed(1),
        from: this.rangeForm.minPrice,
        to: this.rangeForm.maxPrice
      })
      this.loadingTokensInfoStatus = false
    },
    async getTokensBalanceInWallet () {
      console.log('load TokensBalanceInWallet')
      console.log('token0addr=', this.token0Address, 'token1addr=', this.token1Address)
      console.log('token0Decimal=', this.token0Decimal)
      var token0Contract = await this.contractInstance(ViaLendTokenABI, this.token0Address)
      var token1Contract = await this.contractInstance(ViaLendTokenABI, this.token1Address)
      if (this.isConnected) {
        var token0BalanceWeiInWallet = await token0Contract.methods.balanceOf(ethereum.selectedAddress).call()
        this.token0BalanceInWallet = (Number(token0BalanceWeiInWallet) / Number(Math.pow(10, this.token0Decimal))).toFixed(2)
        this.token0BalanceUSDInWallet = Number(this.token0BalanceInWallet) * Number(this.$store.state.token0RateOfUSD)
        // Token1 balance in wallet and vault
        var token1BalanceWeiInWallet = await token1Contract.methods.balanceOf(ethereum.selectedAddress).call()
        this.token1BalanceInWallet = (Number(token1BalanceWeiInWallet) / Number(Math.pow(10, this.token1Decimal))).toFixed(2)
        this.token1BalanceUSDInWallet = Number(this.token1BalanceInWallet) * Number(this.$store.state.token1RateOfUSD)
      }
      var token0BalanceWeiInVault = await token0Contract.methods.balanceOf(this.vaultAddress).call()
      this.token0BalanceInVault = (Number(token0BalanceWeiInVault) / Number(Math.pow(10, this.token0Decimal))).toFixed(2)
      this.token0BalanceUSDInVault = Number(this.token0BalanceInVault) * Number(this.$store.state.token0RateOfUSD)
      console.log('token0BalanceWeiInVault=', this.token0BalanceWeiInVault)
      var token1BalanceWeiInVault = await token1Contract.methods.balanceOf(this.vaultAddress).call()
      this.token1BalanceInVault = (Number(token1BalanceWeiInVault) / Number(Math.pow(10, this.token1Decimal))).toFixed(2)
      this.token1BalanceUSDInVault = Number(this.token1BalanceInVault) * Number(this.$store.state.token1RateOfUSD)
    },
    async loadFees () {
      var accumulateUniswapFees0 = await this.keeperContract.methods.AccumulateUniswapFees0().call()
      var accumulateUniswapFees1 = await this.keeperContract.methods.AccumulateUniswapFees1().call()
      this.accruedProtocolFees0 = await this.keeperContract.methods.accruedProtocolFees0().call()
      this.accruedProtocolFees1 = await this.keeperContract.methods.accruedProtocolFees1().call()
      console.log('accumulateUniswapFees0=', accumulateUniswapFees0)
      console.log('accumulateUniswapFees1=', accumulateUniswapFees1)
      console.log('accruedProtocolFees0=', this.accruedProtocolFees0)
      console.log('accruedProtocolFees1=', this.accruedProtocolFees1)
    },
    minPriceChange (val) {
      // console.log('min price:', val)
      $('.js-range-slider').data('ionRangeSlider').update({ from: val })
    },
    maxPriceChange (val) {
      // console.log('max price:', val)
      $('.js-range-slider').data('ionRangeSlider').update({ to: val })
    },
    drawTickRangeChart () {
      if (this.tickLower !== 0 && this.tickUpper !== 0) {
        var leftMargin = this.tickLower - 10000
        var rightMargin = this.tickUpper + 10000
        var leftPercent = parseInt(((this.tickLower - leftMargin) / (rightMargin - leftMargin)) * 100)
        var rightPercent = parseInt(((this.tickUpper - leftMargin) / (rightMargin - leftMargin)) * 100)
        this.currentPercent = parseInt(((this.currentTick - leftMargin) / (rightMargin - leftMargin)) * 100)
        this.tickRange = [leftPercent, rightPercent]
        console.log('currentPercent=', this.currentPercent)
        if (this.tickLower > this.currentTick || this.tickUpper < this.currentTick) {
          this.rangeStatusStyle = 'outOfRange'
          this.rangeStatus = 'Out of range'
        } else {
          this.rangeStatusStyle = 'inRange'
          this.rangeStatus = 'In range'
        }
        this.rangeStatusDisplay = ''
      }
    },
    async SetUniPortionRatio () {
      var _this = this
      this.uniPortionRatioLoading = true
      if (this.keeperContract != null) {
        this.keeperContract.methods
          .setUniPortionRatio(
            BigInt(this.uniPortionRatio)
          )
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '80000000000',
            // gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // if (showMessage === false) {
            //   _this.$message('Set UniPortionRatio Successful!')
            //   showMessage = true
            // }
          })
          .on('receipt', function (receipt) {
            if (_this.uniPortionRatioLoading) {
              _this.uniPortionRatioLoading = false
            }
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.uniPortionRatioEtherscanDisable = false
              _this.$message('UniPortionRatio submitted!')
              _this.loadTokensBalance()
            } else { _this.$message('UniPortionRatio failed!') }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.uniPortionRatioLoading = false
            console.log(error)
          })
      }
    },
    async SetMaxTotalSupply () {
      var _this = this
      this.maxTotalSupplyLoading = true
      if (this.keeperContract != null) {
        this.keeperContract.methods
          .setMaxTotalSupply(
            BigInt(this.maxTotalSupply)
          )
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '80000000000',
            // gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // if (showMessage === false) {
            //   _this.$message('Set MaxTotalSupply Successful!')
            //   showMessage = true
            // }
          })
          .on('receipt', function (receipt) {
            if (_this.maxTotalSupplyLoading) {
              _this.maxTotalSupplyLoading = false
            }
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.maxTotalSupplyEtherscanDisable = false
              _this.$message('MaxTotalSupply submitted!')
              _this.loadTokensBalance()
            } else { _this.$message('MaxTotalSupply failed!') }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.maxTotalSupplyLoading = false
            console.log(error)
          })
      }
    },
    async SetGovernance () {
      var _this = this
      this.governanceLoading = true
      if (this.keeperContract != null) {
        this.keeperContract.methods
          .setGovernance(
            this.governanceAddress
          )
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '80000000000',
            // gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // if (showMessage === false) {
            //   _this.$message('Set Governance Successful!')
            //   showMessage = true
            // }
          })
          .on('receipt', function (receipt) {
            if (_this.governanceLoading) {
              _this.governanceLoading = false
            }
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.governanceEtherscanDisable = false
              _this.$message('Governance submitted!')
              _this.loadTokensBalance()
            } else { _this.$message('Governance failed!') }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.governanceLoading = false
            console.log(error)
          })
      }
    },
    async acceptGovernance () {
      var _this = this
      this.acceptGovernanceLoading = true
      if (this.keeperContract != null) {
        this.keeperContract.methods
          .acceptGovernance()
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '80000000000',
            // gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // if (showMessage === false) {
            //   _this.$message('Accept Governance Successful!')
            //   showMessage = true
            // }
          })
          .on('receipt', function (receipt) {
            if (_this.acceptGovernanceLoading) {
              _this.acceptGovernanceLoading = false
            }
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.governanceEtherscanDisable = false
              _this.$message('Accept Governance submitted!')
              _this.loadTokensBalance()
            } else { _this.$message('Accept Governance failed!') }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.acceptGovernanceLoading = false
            console.log(error)
          })
      }
    },
    async SetTeam () {
      var _this = this
      this.teamLoading = true
      if (this.keeperContract != null) {
        this.keeperContract.methods
          .setTeam(
            this.teamAddress
          )
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '80000000000',
            // gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // if (showMessage === false) {
            //   _this.$message('SetTeam Successful!')
            //   showMessage = true
            // }
          })
          .on('receipt', function (receipt) {
            if (_this.teamLoading) {
              _this.teamLoading = false
            }
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.teamEtherscanDisable = false
              _this.$message('Team submitted!')
              _this.loadTokensBalance()
            } else { _this.$message('Team failed!') }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.teamLoading = false
            console.log(error)
          })
      }
    },
    async SetProtocolFee () {
      var _this = this
      this.protocolFeeLoading = true
      if (this.keeperContract != null) {
        this.keeperContract.methods
          .setProtocolFee(
            this.protocolFee
          )
          .send({
            from: ethereum.selectedAddress,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
          })
          .on('receipt', function (receipt) {
            if (_this.protocolFeeLoading) {
              _this.protocolFeeLoading = false
            }
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.protocolFeeEtherscanDisable = false
              _this.$message('ProtocolFee submitted!')
              // _this.loadTokensBalance()
            } else { _this.$message('ProtocolFee failed!') }
          })
          .on('error', function (error) {
            _this.$message.error(error)
            _this.protocolFeeLoading = false
            console.log(error)
          })
      }
    },
    async EmergencyBurn () {
      var _this = this
      this.emergencyLoading = true
      if (this.keeperContract != null) {
        this.keeperContract.methods
          .emergencyBurn()
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '80000000000',
            // gas: 600000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // if (showMessage === false) {
            //   _this.$message('Emergency Burn Successful!')
            //   showMessage = true
            // }
          })
          .on('receipt', function (receipt) {
            if (_this.emergencyLoading === true) {
              _this.emergencyLoading = false
            }
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.emergencyEtherscanDisable = false
              _this.$message('Emergency submitted!')
              _this.loadTokensBalance()
            } else { _this.$message('Emergency failed!') }
          })
          .on('error', function (error) {
            _this.emergencyLoading = false
            _this.$message.error(error)
            console.log(error)
          })
      }
    },
    exchangeTokenChange (val) {
      console.log('token change=', val)
      for (var i = 0; i < this.tokensList.length; i++) {
        if (this.tokensList[i].symbol === val) {
          this.usdTokenVal = this.tokensList[i].price_usd
          this.priceUSD = this.tokensList[i].price_usd
          this.currTokenVal = 1
          break
        }
      }
    },
    doRebalance () {
      if (!this.isConnected) {
        this.$message('Please connect wallet!')
        return
      }
      var _this = this
      this.rebalanceLoading = true
      if (isNaN(this.rangeForm.minPrice) || this.rangeForm.minPrice <= 0) {
        this.$message({
          message: 'Please input positive number greater than 0 in Min Price.',
          type: 'warning'
        })
      }
      if (isNaN(this.rangeForm.maxPrice) || this.rangeForm.maxPrice <= 0) {
        this.$message({
          message: 'Please input positive number greater than 0 in Max Price.',
          type: 'warning'
        })
      }
      if (this.keeperContract != null) {
        console.log('account address is ' + ethereum.selectedAddress)
        console.log('this.rangeForm.tickLower=', Math.round(parseInt(this.tickLower) / 60) * 60)
        console.log('this.rangeForm.tickUpper=', Math.round(parseInt(this.tickUpper) / 60) * 60)

        this.keeperContract.methods
          .strategy1(
            BigInt(Math.round(parseInt(this.tickLower) / 60) * 60),
            BigInt(Math.round(parseInt(this.tickUpper) / 60) * 60)
          )
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '1000000000',
            // gas: 900000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
            // if (_this.rebalanceLoading) {
            //   _this.rebalanceLoading = false
            //   _this.$message('rebalance confirmed!')
            //   console.log('confirmation,confirmationNumber:' + confirmationNumber + ',receipt:' + JSON.stringify(receipt))
            //   if (receipt != null) {
            //     _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
            //     _this.viewOnEtherscanDisable = false
            //   }
            //   // _this.errorRebalance = receipt
            //   _this.initData()
            // }
          })
          .on('receipt', function (receipt) {
            if (_this.rebalanceLoading) {
              _this.rebalanceLoading = false
            }
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.doRebalanceEtherscanDisable = false
              _this.$message('Rebalance submitted!')
              _this.loadTokensBalance()
            } else { _this.$message('Rebalance failed!') }
          })
          .on('error', function (error) {
            _this.rebalanceLoading = false
            _this.$message.error(error)
          })
      }
    },
    doRemovePositions () {
      if (!this.isConnected) {
        this.$message('Please connect wallet!')
        return
      }
      var _this = this
      this.removePositionsLoading = true
      if (this.keeperContract != null) {
        this.keeperContract.methods
          .alloc()
          .send({
            from: ethereum.selectedAddress,
            // gasPrice: '1000000000',
            // gas: 900000,
            value: 0
          })
          .on('confirmation', function (confirmationNumber, receipt) {
          })
          .on('receipt', function (receipt) {
            if (_this.removePositionsLoading) {
              _this.removePositionsLoading = false
            }
            if (receipt.status) {
              _this.goToEtherscan = 'https://goerli.etherscan.io/tx/' + receipt.transactionHash
              _this.removePositionsEtherscanDisable = false
              _this.$message('RemovePositions submitted!')
              _this.loadTokensBalance()
            } else { _this.$message('RemovePositions failed!') }
          })
          .on('error', function (error) {
            _this.removePositionsLoading = false
            _this.$message.error(error)
          })
      }
    },
    priceToTick (price, decimal0, decimal1) {
      console.log('price=', price, 'decimal0=', decimal0, 'decimal1=', decimal1)
      // return parseInt(Math.log(price / Math.pow(10, (decimal0 - decimal1))) / Math.log(1.0001))
      return Math.log(price / Math.pow(10, (decimal0 - decimal1))) / Math.log(1.0001)
    },
    tickToPrice (tick, decimal0, decimal1) {
      console.log('decimal0=', decimal0, 'decimal1=', decimal1)
      console.log('Math.pow(1.0001, tick)=', Math.pow(1.0001, tick), 'decimal=', decimal0 - decimal1)
      // var price = Math.pow(1.0001, tick) * Math.pow(10, (decimal0 - decimal1))
      // if (price.toFixed(1) > 0) {
      //   this.precision = 1
      //   this.step = 0.1
      //   return price.toFixed(1)
      // } else {
      //   this.precision = 3
      //   this.step = 0.001
      //   return price.toFixed(3)
      // }
      return Math.pow(1.0001, tick) * Math.pow(10, (decimal0 - decimal1))
    },
    onSubmit () {
      console.log('submit!')
    },
    get_liquidity_0 (x, sa, sb) {
      return x * sa * sb / (sb - sa)
    },
    calculate_y (L, sp, sa) {
      return L * (sp - sa)
    },
    calculate_c (p, d, x, y) {
      return y / ((d - 1) * p * x + y)
    },
    calculate_d (p, c, x, y) {
      return 1 + y * (1 - c) / (c * p * x)
    },
    getY () {
      var p = 3879.97
      var a = 300
      var b = 5000
      var x = 1
      var sp = Math.sqrt(p)
      var sa = Math.sqrt(a)
      var sb = Math.sqrt(b)
      console.log('Math.pow(p,0.5)=', Math.pow(p, 0.5), ';sp=', sp)

      var L = this.get_liquidity_0(x, sp, sb)
      var y = this.calculate_y(L, sp, sa)
      console.log('amount of USDC y=', y)
      var c = sb / sp
      var d = sa / sp
      var ic = (p, d, x, y)
      console.log('ic=,', ic)

      var id = this.calculate_d(p, c, x, y)
      console.log('id=', id)
      var C = Math.pow(ic, 2)
      console.log('C=', C)
      var D = Math.pow(id, 2)
      console.log('D=', D)
      console.log('p_a=', D * p, ',(', D * 100, ' of P),p_b=', C * p, ',(', C * 100, ' of P')
    },
    async getPairsSettingData () {
      this.maxTotalSupply = await this.keeperContract.methods.maxTotalSupply().call()
      this.governanceAddress = await this.keeperContract.methods.governance().call()
      this.teamAddress = await this.keeperContract.methods.team().call()
      this.uniPortionRatio = await this.keeperContract.methods.uniPortion().call()
      console.log('uniPortionRatio=', this.uniPortionRatio)
      this.protocolFee = await this.keeperContract.methods.protocolFee().call()
    },
    async getPairPublicInfo (pair) {
      this.tvlDataLoading = true
      this.assetsRatioLoading = true
      console.log('pair.vaultAddress=', pair.vaultAddress)
      console.log('pair.token0LendingAddress=', pair.token0.token0LendingAddress)
      console.log('pair.token1LendingAddress=', pair.token1.token1LendingAddress)
      console.log('pair.token0.tokenAddress=', pair.token0.tokenAddress)
      console.log('pair.token1.tokenAddress=', pair.token1.tokenAddress)
      var keeperContract = await this.contractInstance(contractABI, pair.vaultAddress)
      var token0LendingContract = await this.contractInstance(ViaLendTokenABI, pair.token0.token0LendingAddress)
      var token1LendingContract = await this.contractInstance(ViaLendTokenABI, pair.token1.token1LendingAddress)
      var token0Contract = await this.contractInstance(ViaLendTokenABI, pair.token0.tokenAddress)
      var token1Contract = await this.contractInstance(ViaLendTokenABI, pair.token1.tokenAddress)
      // ---------- Get TVL Begin-----------------
      var uniliqs = await keeperContract.methods.getPositionAmounts(BigInt(pair.cLow), BigInt(pair.cHigh)).call()
      console.log('balance in uniswap:', uniliqs, 'getVaultInfo_cLow=', pair.cLow, 'getVaultInfo_cHigh=', pair.cHigh)
      // -->Get Lending Amounts begin
      var exchangeRate0 = await token0LendingContract.methods.exchangeRateStored().call()
      var exchangeRate1 = await token1LendingContract.methods.exchangeRateStored().call()
      var CAmount0 = await token0LendingContract.methods.balanceOf(pair.vaultAddress).call()
      var CAmount1 = await token1LendingContract.methods.balanceOf(pair.vaultAddress).call()
      var underlying0 = CAmount0 * exchangeRate0 / Math.pow(10, 18)
      var underlying1 = CAmount1 * exchangeRate1 / Math.pow(10, 18)
      console.log('underlying0=', underlying0, 'underlying1=', underlying1)
      // -->Get Lending Amounts end
      pair.vaultLending = Number(underlying0) + Number(underlying1) // Not yet converted to USD
      pair.token0BalanceInVault = await token0Contract.methods.balanceOf(pair.vaultAddress).call()
      pair.token1BalanceInVault = await token1Contract.methods.balanceOf(pair.vaultAddress).call()
      console.log('token0BalanceInVault=', pair.token0BalanceInVault, ';token1BalanceInVault=', pair.token1BalanceInVault)
      var t0Decimal = await token0Contract.methods.decimals().call()
      var t1Decimal = await token1Contract.methods.decimals().call()
      pair.tvlTotal0 = (Number(pair.token0BalanceInVault) + Number(uniliqs.amount0) + Number(underlying0)) / Number(Math.pow(10, t0Decimal))
      pair.tvlTotal1 = (Number(pair.token1BalanceInVault) + Number(uniliqs.amount1) + Number(underlying1)) / Number(Math.pow(10, t1Decimal))
      // pair.tvl = pair.tvlTotal0 * 3768.67 + pair.tvlTotal1 * 0.998117
      console.log('tvlTotal0=', pair.tvlTotal0, 'tvlTotal1=', pair.tvlTotal1)
      // console.log('TVL=', this.tvl, ';token0RateOfUSD=', this.pairsList[0].smartVaults[0].rateOfUSD, ';token1RateOfUSD=', this.pairsList[0].smartVaults[1].rateOfUSD)
      // ---------- Get TVL End--------------------------
      // ---------- Get Assets ratio Begin --------------
      var result = await keeperContract.methods.getPositionAmounts(BigInt(pair.cLow), BigInt(pair.cHigh)).call()
      var token0BalanceInPool, token1BalanceInPool, token0BalanceInLending, token1BalanceInLending
      if (result !== undefined && result !== null) {
        token0BalanceInPool = result.amount0
        token1BalanceInPool = result.amount1
      }
      if (!isNaN(underlying0) && !isNaN(underlying1)) {
        token0BalanceInLending = underlying0
        token1BalanceInLending = underlying1
      }
      var totalUniswap = Number(token0BalanceInPool) * 300 + Number(token1BalanceInPool)
      var totalLending = Number(token0BalanceInLending) * 300 + Number(token1BalanceInLending)
      var totalUsdc = totalUniswap + totalLending
      if (totalUsdc === 0) {
        pair.uniswapRatio = 0
        pair.lendingRatio = 0
      } else {
        pair.uniswapRatio = totalUniswap / totalUsdc * 100
        pair.lendingRatio = totalLending / totalUsdc * 100
      }
      if ((Number(token0BalanceInPool) + Number(token0BalanceInLending)) === 0) {
        pair.uniToken0Rate = 0
      } else {
        pair.uniToken0Rate = Number(token0BalanceInPool) / (Number(token0BalanceInPool) + Number(token0BalanceInLending))
      }
      if ((Number(token1BalanceInPool) + Number(token1BalanceInLending)) === 0) {
        pair.uniToken1Rate = 0
      } else {
        pair.uniToken1Rate = Number(token1BalanceInPool) / (Number(token1BalanceInPool) + Number(token1BalanceInLending))
      }
      if ((Number(token0BalanceInPool) + Number(token0BalanceInLending)) === 0) {
        pair.lendingToken0Rate = 0
      } else {
        pair.lendingToken0Rate = Number(token0BalanceInLending) / (Number(token0BalanceInPool) + Number(token0BalanceInLending))
      }
      if ((Number(token1BalanceInPool) + Number(token1BalanceInLending)) === 0) {
        pair.lendingToken1Rate = 0
      } else {
        pair.lendingToken1Rate = Number(token1BalanceInLending) / (Number(token1BalanceInPool) + Number(token1BalanceInLending))
      }
      console.log('totalUniswap=', pair.totalUniswap, 'totalLending=', pair.totalLending, 'total_usdc=', pair.totalUsdc, 'uniswapRatio=', pair.uniswapRatio, 'lendingRatio=', pair.lendingRatio)
      // ---------- Get Assets ratio End --------------
      console.log('ethereum.selectedAddress=', ethereum.selectedAddress)
      if (ethereum.selectedAddress !== null && ethereum.selectedAddress !== undefined) {
        var myShare, totalShares, tvlTotal0, tvlTotal1
        myShare = await keeperContract.methods.balanceOf(ethereum.selectedAddress).call()
        totalShares = await keeperContract.methods.totalSupply().call()
        var Assets = await keeperContract.methods.Assetholder(ethereum.selectedAddress).call()
        if (Assets !== null) {
          // calc APR
          var oraclePriceTwap
          var poolAddress = await keeperContract.methods.poolAddress().call()
          var twapDuration = 2
          // this.sleep(3000)
          var keeperUniswapV3Contract = this.contractInstance(uniswapV3PoolABI, poolAddress)
          var slot0 = await keeperUniswapV3Contract.methods.slot0().call()
          if (slot0 !== null && slot0 !== undefined) {
            var twap = slot0['tick']
            oraclePriceTwap = await keeperContract.methods.getQuoteAtTick(Number(twap), BigInt(Math.pow(10, 18)), pair.token0.tokenAddress, pair.token1.tokenAddress).call()
            console.log('twap=', twap, ';oraclePriceTwap=', oraclePriceTwap)
          }
          var Total0, Total1

          Total0 = Number(pair.token0BalanceInVault) + Number(uniliqs.amount0) + Number(underlying0)
          Total1 = Number(pair.token1BalanceInVault) + Number(uniliqs.amount1) + Number(underlying1)
          var mytvl0, mytvl1
          if (Number(totalShares) === 0) {
            mytvl0 = 0
            mytvl1 = 0
          } else {
            mytvl0 = Total0 * Number(myShare) / Number(totalShares)
            mytvl1 = Total1 * Number(myShare) / Number(totalShares)
          }

          var oneyearINsec = 365 * 24 * 60 * 60
          var block = await web3.eth.getBlock(Assets.block)
          console.log('block timestamp=', block.timestamp)
          var timestamp = block.timestamp
          var headerNumber = await web3.eth.getBlockNumber()
          var headerBlock = await web3.eth.getBlock(headerNumber)
          // console.log('header=', JSON.stringify(headerBlock))
          var htimestamp = headerBlock.timestamp
          console.log('htimestamp=', htimestamp, ';timestamp=', timestamp)
          var timediff = Number(htimestamp) - Number(timestamp)
          var deposit0 = Assets.deposit0
          var deposit1 = Assets.deposit1
          var fd0 = Number(deposit0)
          var fd1, fm1
          if (oraclePriceTwap === 0) {
            fd1 = 0
            fm1 = 0
          } else {
            fd1 = Number(deposit1) * Math.pow(10, Number(pair.token0.decimals)) / oraclePriceTwap
            fm1 = Number(mytvl1) * Math.pow(10, Number(pair.token0.decimals)) / oraclePriceTwap
          }
          var fm0 = Number(mytvl0)
          var fdd = fd0 + fd1
          var fmm = fm0 + fm1
          console.log('fm0=', fm0, 'fm1=', fm1)
          console.log('decimal=', pair.token0.decimals)
          console.log('fmm=', fmm, 'fdd=', fdd, 'timediff=', timediff, 'oneyearInsec=', oneyearINsec)
          if (fdd === 0 || timediff === 0) {
            pair.currentAPR = 0
          } else {
            pair.currentAPR = (fmm - fdd) / Number(timediff) * oneyearINsec / fdd
          }
        }
      }
      return pair
    }

  }
}
</script>

<style scoped>
.el-header {
  padding: 0px !important;
}
.el-main {
  padding: 0px !important;
}
.el-menu {
  border-right: 0px !important;
}
.el-aside {
  background-color: #304156 !important;
}
.breadcrumb {
  padding: 20px;
  margin-bottom: 0px !important;
}
.block {
  float: left;
}
.el-divider--horizontal {
  margin: 0px !important;
}
.range_title {
  font-size: 20px;
  color: black;
  padding: 10px;
}
.token_exchange_form {
  width: 100%;
  font-size: 20px;
  color: black;
  padding: 10px;
}
.dot {
  position: absolute;
  width: 10px;
  height: 10px;
  border-radius: 100%;
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.2);
}
.status {
  margin-left: 15px;
}
.inRange {
  background: #0d7404;
}
.outOfRange {
  background: brown;
}
.list-group li {
  padding: 20px;
  color: white;
}
.list-group a {
  color: white;
}
.list-group-item {
  background-color: #545c64 !important;
}
.list-group-item.active,
.list-group-item.active:focus,
.list-group-item.active:hover {
  background-color: #337ab7 !important;
  border-color: #337ab7 !important;
}
.menu_text a {
  color: #fff;
}
.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: '';
}
.clearfix:after {
  clear: both;
}

.box-card {
  width: 280px;
  margin: 10px;
}
.box-card-currency {
  width: 900px;
  margin: 10px;
}
.el-form-item__label {
  line-height: 20px !important;
}
html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
}
#app,
.box {
  height: 100%;
}
.el-container {
  height: 100%;
}
.el-input-number.is-without-controls .el-input__inner {
  width: 200px;
  line-height: 30px;
  height: 28px;
}
</style>

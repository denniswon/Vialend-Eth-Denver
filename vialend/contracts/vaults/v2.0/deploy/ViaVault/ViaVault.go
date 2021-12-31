// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vaultCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_individualCap\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Assetholder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"current1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn1\",\"type\":\"uint256\"}],\"name\":\"calcPositionShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"p\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a1\",\"type\":\"uint256\"}],\"name\":\"calcShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"checkCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn1\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getbalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getbalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"individualCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moveFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setIndividualCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strat\",\"type\":\"address\"}],\"name\":\"setSrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setVaultCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"percent\",\"type\":\"uint8\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x610100604052600c805460ff191690553480156200001c57600080fd5b506040516200251a3803806200251a8339810160408190526200003f91620001b0565b604080518082018252601a81527f5669614c656e6420556e6920436f6d706f756e6420546f6b656e0000000000006020808301918252835180850190945260048452630565543360e41b908401528151919291620000a091600391620000ed565b508051620000b6906004906020840190620000ed565b50506001600555506001600160a01b0395861660805293851660a05291841660c05290921660e052600a91909155600b556200025c565b828054620000fb906200021f565b90600052602060002090601f0160209004810192826200011f57600085556200016a565b82601f106200013a57805160ff19168380011785556200016a565b828001600101855582156200016a579182015b828111156200016a5782518255916020019190600101906200014d565b50620001789291506200017c565b5090565b5b808211156200017857600081556001016200017d565b80516001600160a01b0381168114620001ab57600080fd5b919050565b60008060008060008060c08789031215620001ca57600080fd5b620001d58762000193565b9550620001e56020880162000193565b9450620001f56040880162000193565b9350620002056060880162000193565b92506080870151915060a087015190509295509295509295565b600181811c908216806200023457607f821691505b602082108114156200025657634e487b7160e01b600052602260045260246000fd5b50919050565b60805160a05160c05160e051612214620003066000396000818161062b01528181610a9801528181610c070152818161111d015281816112e2015261163201526000818161028601528181610bca015281816110e9015281816111fa015281816112a201526115f701526000818161073a01528181610b58015281816111a0015281816113ab01526116b901526000818161057801528181611327015261146701526122146000f3fe6080604052600436106101f05760003560e01c8063a457c2d71161010c578063d0b6ba2a1161009a578063dd62ed3e1161006c578063dd62ed3e14610682578063e2bbb158146106c8578063f0f5907d146106e8578063f2a40db814610708578063f851a4401461072857005b8063d0b6ba2a14610604578063d21220a714610619578063d46915bc1461064d578063d8d7f96f1461066d57005b8063b49c308e116100de578063b49c308e14610546578063c45a015514610566578063c6ab5d901461059a578063cd2b3026146105cf578063ceadeae0146105ef57005b8063a457c2d7146104b9578063a8c62e76146104d9578063a9059cbb146104f9578063aabfd5721461051957005b806333e2f2ea116101895780635bb6aa851161015b5780635bb6aa85146103b75780635ccaf589146104345780636ea056a91461044e57806370a082311461046e57806395d89b41146104a457005b806333e2f2ea14610331578063385602f81461036c57806339509351146103815780633c1bda09146103a157005b806318160ddd116101c257806318160ddd146102c05780631ba326c4146102d557806323b872dd146102f5578063313ce5671461031557005b80630276650b146101f957806306fdde0314610222578063095ea7b3146102445780630dfe16811461027457005b366101f757005b005b34801561020557600080fd5b5061020f600b5481565b6040519081526020015b60405180910390f35b34801561022e57600080fd5b5061023761075c565b6040516102199190611ef6565b34801561025057600080fd5b5061026461025f366004611f45565b6107ee565b6040519015158152602001610219565b34801561028057600080fd5b506102a87f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610219565b3480156102cc57600080fd5b5060025461020f565b3480156102e157600080fd5b5061020f6102f0366004611f6f565b610804565b34801561030157600080fd5b50610264610310366004611f9b565b610824565b34801561032157600080fd5b5060405160128152602001610219565b34801561033d57600080fd5b5061035161034c366004611fd7565b6108d3565b60408051938452602084019290925290820152606001610219565b34801561037857600080fd5b5061020f610a80565b34801561038d57600080fd5b5061026461039c366004611f45565b610b11565b3480156103ad57600080fd5b5061020f600a5481565b3480156103c357600080fd5b506104076103d2366004611ff9565b600860205260009081526040902080546001820154600283015460038401546004850154600590950154939492939192909186565b604080519687526020870195909552938501929092526060840152608083015260a082015260c001610219565b34801561044057600080fd5b50600c546102649060ff1681565b34801561045a57600080fd5b506101f7610469366004611f45565b610b4d565b34801561047a57600080fd5b5061020f610489366004611ff9565b6001600160a01b031660009081526020819052604090205490565b3480156104b057600080fd5b50610237610c86565b3480156104c557600080fd5b506102646104d4366004611f45565b610c95565b3480156104e557600080fd5b506009546102a8906001600160a01b031681565b34801561050557600080fd5b50610264610514366004611f45565b610d2e565b34801561052557600080fd5b5061020f610534366004611ff9565b60076020526000908152604090205481565b34801561055257600080fd5b5061020f610561366004611fd7565b610d3b565b34801561057257600080fd5b506102a87f000000000000000000000000000000000000000000000000000000000000000081565b3480156105a657600080fd5b506105ba6105b5366004612014565b610dd7565b60408051928352602083019190915201610219565b3480156105db57600080fd5b506101f76105ea366004612037565b611195565b3480156105fb57600080fd5b5061020f6111e2565b34801561061057600080fd5b506101f7611231565b34801561062557600080fd5b506102a87f000000000000000000000000000000000000000000000000000000000000000081565b34801561065957600080fd5b506101f7610668366004611ff9565b61131c565b34801561067957600080fd5b506101f76113a0565b34801561068e57600080fd5b5061020f61069d366004612050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3480156106d457600080fd5b506101f76106e3366004611fd7565b611452565b3480156106f457600080fd5b506101f7610703366004612037565b6116ae565b34801561071457600080fd5b506102a8610723366004612037565b6116fb565b34801561073457600080fd5b506102a87f000000000000000000000000000000000000000000000000000000000000000081565b60606003805461076b90612083565b80601f016020809104026020016040519081016040528092919081815260200182805461079790612083565b80156107e45780601f106107b9576101008083540402835291602001916107e4565b820191906000526020600020905b8154815290600101906020018083116107c757829003601f168201915b5050505050905090565b60006107fb338484611725565b50600192915050565b600061081084836120d4565b61081a90846120f6565b90505b9392505050565b6000610831848484611849565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156108bb5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b6108c88533858403611725565b506001949350505050565b60095460408051634c6afee560e11b815290516000928592859285926001600160a01b0316916398d5fdca9160048083019260209291908290030181865afa158015610923573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610947919061210e565b90506000811161097e5760405162461bcd60e51b8152602060048201526002602482015261070360f41b60448201526064016108b2565b600061098960025490565b9050806109a25761099b828585610804565b9450610a77565b60095460408051636253bb0f60e11b8152815160009384936001600160a01b039091169263c4a7761e9260048083019392829003018187875af11580156109ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a119190612127565b91509150610a1d6111e2565b610a2790836120f6565b9150610a31610a80565b610a3b90826120f6565b90506000610a4a858484610804565b90506000610a59868c8c610804565b905081610a66868361214b565b610a7091906120d4565b9850505050505b50509250925092565b6040516370a0823160e01b81523060048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a08231906024015b602060405180830381865afa158015610ae8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0c919061210e565b905090565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916107fb918590610b489086906120f6565b611725565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b955760405162461bcd60e51b81526004016108b29061216a565b33610bc85760405162461bcd60e51b81526020600482015260036024820152620e660f60eb1b60448201526064016108b2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614158015610c3c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614155b610c6e5760405162461bcd60e51b8152602060048201526003602482015262746f6b60e81b60448201526064016108b2565b610c826001600160a01b0383163383611a17565b5050565b60606004805461076b90612083565b3360009081526001602090815260408083206001600160a01b038616845290915281205482811015610d175760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016108b2565b610d243385858403611725565b5060019392505050565b60006107fb338484611849565b600080600960009054906101000a90046001600160a01b03166001600160a01b03166398d5fdca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db5919061210e565b9050600083610dc4838761214b565b610dce91906120f6565b95945050505050565b60008033610e0d5760405162461bcd60e51b81526020600482015260036024820152620ee60f60eb1b60448201526064016108b2565b60648360ff161115610e465760405162461bcd60e51b8152602060048201526002602482015261706360f01b60448201526064016108b2565b33600090815260208190526040812054606490610e679060ff87169061214b565b610e7191906120d4565b90506000610e7e60025490565b9050600082118015610e905750808211155b610ec55760405162461bcd60e51b815260206004820152600660248201526573686172657360d01b60448201526064016108b2565b610ecf3383611a7f565b60095460408051636253bb0f60e11b8152815160009384936001600160a01b039091169263c4a7761e9260048083019392829003018187875af1158015610f1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3e9190612127565b91509150600080610f4d6111e2565b610f55610a80565b9092509050600080610f6784876120f6565b610f7184876120f6565b909250905086610f81898461214b565b610f8b91906120d4565b995086610f98898361214b565b610fa291906120d4565b985060008a851080610fb357508984105b1561109e576009546040516371c851dd60e01b8152336004820152602481018d9052604481018c90526001600160a01b03909116906371c851dd906064016020604051808303816000875af1158015611010573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110349190612189565b90506000806110416111e2565b611049610a80565b915091508c821015801561105d57508b8110155b6110975760405162461bcd60e51b815260206004820152600b60248201526a6e65772062616c616e636560a81b60448201526064016108b2565b50506110a2565b5060015b806110dc5760405162461bcd60e51b815260206004820152600a6024820152691cdd1c985d08199d5b9960b21b60448201526064016108b2565b6111106001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016338d611a17565b6111446001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016338c611a17565b604080518a8152602081018d90529081018b905233907f02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca949060600160405180910390a2505050505050505050915091565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146111dd5760405162461bcd60e51b81526004016108b29061216a565b600b55565b6040516370a0823160e01b81523060048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401610acb565b6009546001600160a01b031633146112735760405162461bcd60e51b81526020600482015260056024820152641cdd1c985d60da1b60448201526064016108b2565b600061127d6111e2565b90506000611289610a80565b905081156112cb576009546112cb906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116911684611a17565b801561130b5760095461130b906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116911683611a17565b5050600c805460ff19166001179055565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461137e5760405162461bcd60e51b8152602060048201526007602482015266666163746f727960c81b60448201526064016108b2565b600980546001600160a01b0319166001600160a01b0392909216919091179055565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146113e85760405162461bcd60e51b81526004016108b29061216a565b600960009054906101000a90046001600160a01b03166001600160a01b0316637c7311fe6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561143857600080fd5b505af115801561144c573d6000803e3d6000fd5b50505050565b6040516376beea2760e11b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ed7dd44e90602401602060405180830381865afa1580156114b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114da9190612189565b6115135760405162461bcd60e51b815260206004820152600a6024820152696e6f742061637469766560b01b60448201526064016108b2565b60008211806115225750600081115b6115575760405162461bcd60e51b81526004016108b2906020808252600490820152630616d74360e41b604082015260600190565b3361158a5760405162461bcd60e51b81526020600482015260036024820152620c860f60eb1b60448201526064016108b2565b600080600061159985856108d3565b925092509250600083116115da5760405162461bcd60e51b8152602060048201526008602482015267073686172653c3d360c41b60448201526064016108b2565b6115e43384611bcd565b811561161f5761161f6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333085611cac565b801561165a5761165a6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333084611cac565b600c805460ff19169055604080518481526020810184905290810182905233907f36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e9060600160405180910390a25050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146116f65760405162461bcd60e51b81526004016108b29061216a565b600a55565b6006818154811061170b57600080fd5b6000918252602090912001546001600160a01b0316905081565b6001600160a01b0383166117875760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016108b2565b6001600160a01b0382166117e85760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016108b2565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b0383166118ad5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016108b2565b6001600160a01b03821661190f5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016108b2565b6001600160a01b038316600090815260208190526040902054818110156119875760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016108b2565b6001600160a01b038085166000908152602081905260408082208585039055918516815290812080548492906119be9084906120f6565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611a0a91815260200190565b60405180910390a361144c565b6040516001600160a01b038316602482015260448101829052611a7a90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611ce4565b505050565b6001600160a01b038216611adf5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016108b2565b6001600160a01b03821660009081526020819052604090205481811015611b535760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016108b2565b6001600160a01b0383166000908152602081905260408120838303905560028054849290611b829084906121ab565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050565b6001600160a01b038216611c235760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016108b2565b8060026000828254611c3591906120f6565b90915550506001600160a01b03821660009081526020819052604081208054839290611c629084906120f6565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6040516001600160a01b038085166024830152831660448201526064810182905261144c9085906323b872dd60e01b90608401611a43565b6000611d39826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611db69092919063ffffffff16565b805190915015611a7a5780806020019051810190611d579190612189565b611a7a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016108b2565b606061081a8484600085856001600160a01b0385163b611e185760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016108b2565b600080866001600160a01b03168587604051611e3491906121c2565b60006040518083038185875af1925050503d8060008114611e71576040519150601f19603f3d011682016040523d82523d6000602084013e611e76565b606091505b5091509150611e86828286611e91565b979650505050505050565b60608315611ea057508161081d565b825115611eb05782518084602001fd5b8160405162461bcd60e51b81526004016108b29190611ef6565b60005b83811015611ee5578181015183820152602001611ecd565b8381111561144c5750506000910152565b6020815260008251806020840152611f15816040850160208701611eca565b601f01601f19169190910160400192915050565b80356001600160a01b0381168114611f4057600080fd5b919050565b60008060408385031215611f5857600080fd5b611f6183611f29565b946020939093013593505050565b600080600060608486031215611f8457600080fd5b505081359360208301359350604090920135919050565b600080600060608486031215611fb057600080fd5b611fb984611f29565b9250611fc760208501611f29565b9150604084013590509250925092565b60008060408385031215611fea57600080fd5b50508035926020909101359150565b60006020828403121561200b57600080fd5b61081d82611f29565b60006020828403121561202657600080fd5b813560ff8116811461081d57600080fd5b60006020828403121561204957600080fd5b5035919050565b6000806040838503121561206357600080fd5b61206c83611f29565b915061207a60208401611f29565b90509250929050565b600181811c9082168061209757607f821691505b602082108114156120b857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b6000826120f157634e487b7160e01b600052601260045260246000fd5b500490565b60008219821115612109576121096120be565b500190565b60006020828403121561212057600080fd5b5051919050565b6000806040838503121561213a57600080fd5b505080516020909101519092909150565b6000816000190483118215151615612165576121656120be565b500290565b60208082526005908201526430b236b4b760d91b604082015260600190565b60006020828403121561219b57600080fd5b8151801515811461081d57600080fd5b6000828210156121bd576121bd6120be565b500390565b600082516121d4818460208701611eca565b919091019291505056fea2646970667358221220c2086f6306ba684725d9d7049be7e10852a900a162de3b9d29419d6adb138dd364736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _admin common.Address, _token0 common.Address, _token1 common.Address, _vaultCap *big.Int, _individualCap *big.Int) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _factory, _admin, _token0, _token1, _vaultCap, _individualCap)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 current0, uint256 current1, uint256 block, uint256 withdrawShares)
func (_Api *ApiCaller) Assetholder(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit0       *big.Int
	Deposit1       *big.Int
	Current0       *big.Int
	Current1       *big.Int
	Block          *big.Int
	WithdrawShares *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Assetholder", arg0)

	outstruct := new(struct {
		Deposit0       *big.Int
		Deposit1       *big.Int
		Current0       *big.Int
		Current1       *big.Int
		Block          *big.Int
		WithdrawShares *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Deposit1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Current0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Current1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.WithdrawShares = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 current0, uint256 current1, uint256 block, uint256 withdrawShares)
func (_Api *ApiSession) Assetholder(arg0 common.Address) (struct {
	Deposit0       *big.Int
	Deposit1       *big.Int
	Current0       *big.Int
	Current1       *big.Int
	Block          *big.Int
	WithdrawShares *big.Int
}, error) {
	return _Api.Contract.Assetholder(&_Api.CallOpts, arg0)
}

// Assetholder is a free data retrieval call binding the contract method 0x5bb6aa85.
//
// Solidity: function Assetholder(address ) view returns(uint256 deposit0, uint256 deposit1, uint256 current0, uint256 current1, uint256 block, uint256 withdrawShares)
func (_Api *ApiCallerSession) Assetholder(arg0 common.Address) (struct {
	Deposit0       *big.Int
	Deposit1       *big.Int
	Current0       *big.Int
	Current1       *big.Int
	Block          *big.Int
	WithdrawShares *big.Int
}, error) {
	return _Api.Contract.Assetholder(&_Api.CallOpts, arg0)
}

// AccId is a free data retrieval call binding the contract method 0xaabfd572.
//
// Solidity: function accId(address ) view returns(uint256)
func (_Api *ApiCaller) AccId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "accId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccId is a free data retrieval call binding the contract method 0xaabfd572.
//
// Solidity: function accId(address ) view returns(uint256)
func (_Api *ApiSession) AccId(arg0 common.Address) (*big.Int, error) {
	return _Api.Contract.AccId(&_Api.CallOpts, arg0)
}

// AccId is a free data retrieval call binding the contract method 0xaabfd572.
//
// Solidity: function accId(address ) view returns(uint256)
func (_Api *ApiCallerSession) AccId(arg0 common.Address) (*big.Int, error) {
	return _Api.Contract.AccId(&_Api.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Api *ApiCaller) Accounts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "accounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Api *ApiSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Accounts(&_Api.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(address)
func (_Api *ApiCallerSession) Accounts(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Accounts(&_Api.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCallerSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Api *ApiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Api *ApiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Api.Contract.Allowance(&_Api.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Api *ApiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Api.Contract.Allowance(&_Api.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Api *ApiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Api *ApiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Api *ApiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, account)
}

// CalcShare is a free data retrieval call binding the contract method 0x1ba326c4.
//
// Solidity: function calcShare(uint256 p, uint256 a0, uint256 a1) pure returns(uint256)
func (_Api *ApiCaller) CalcShare(opts *bind.CallOpts, p *big.Int, a0 *big.Int, a1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "calcShare", p, a0, a1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcShare is a free data retrieval call binding the contract method 0x1ba326c4.
//
// Solidity: function calcShare(uint256 p, uint256 a0, uint256 a1) pure returns(uint256)
func (_Api *ApiSession) CalcShare(p *big.Int, a0 *big.Int, a1 *big.Int) (*big.Int, error) {
	return _Api.Contract.CalcShare(&_Api.CallOpts, p, a0, a1)
}

// CalcShare is a free data retrieval call binding the contract method 0x1ba326c4.
//
// Solidity: function calcShare(uint256 p, uint256 a0, uint256 a1) pure returns(uint256)
func (_Api *ApiCallerSession) CalcShare(p *big.Int, a0 *big.Int, a1 *big.Int) (*big.Int, error) {
	return _Api.Contract.CalcShare(&_Api.CallOpts, p, a0, a1)
}

// CheckCap is a free data retrieval call binding the contract method 0xb49c308e.
//
// Solidity: function checkCap(uint256 amount0, uint256 amount1) view returns(uint256)
func (_Api *ApiCaller) CheckCap(opts *bind.CallOpts, amount0 *big.Int, amount1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "checkCap", amount0, amount1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckCap is a free data retrieval call binding the contract method 0xb49c308e.
//
// Solidity: function checkCap(uint256 amount0, uint256 amount1) view returns(uint256)
func (_Api *ApiSession) CheckCap(amount0 *big.Int, amount1 *big.Int) (*big.Int, error) {
	return _Api.Contract.CheckCap(&_Api.CallOpts, amount0, amount1)
}

// CheckCap is a free data retrieval call binding the contract method 0xb49c308e.
//
// Solidity: function checkCap(uint256 amount0, uint256 amount1) view returns(uint256)
func (_Api *ApiCallerSession) CheckCap(amount0 *big.Int, amount1 *big.Int) (*big.Int, error) {
	return _Api.Contract.CheckCap(&_Api.CallOpts, amount0, amount1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiSession) Decimals() (uint8, error) {
	return _Api.Contract.Decimals(&_Api.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiCallerSession) Decimals() (uint8, error) {
	return _Api.Contract.Decimals(&_Api.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Api *ApiCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Api *ApiSession) Factory() (common.Address, error) {
	return _Api.Contract.Factory(&_Api.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Api *ApiCallerSession) Factory() (common.Address, error) {
	return _Api.Contract.Factory(&_Api.CallOpts)
}

// Getbalance0 is a free data retrieval call binding the contract method 0xceadeae0.
//
// Solidity: function getbalance0() view returns(uint256)
func (_Api *ApiCaller) Getbalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getbalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Getbalance0 is a free data retrieval call binding the contract method 0xceadeae0.
//
// Solidity: function getbalance0() view returns(uint256)
func (_Api *ApiSession) Getbalance0() (*big.Int, error) {
	return _Api.Contract.Getbalance0(&_Api.CallOpts)
}

// Getbalance0 is a free data retrieval call binding the contract method 0xceadeae0.
//
// Solidity: function getbalance0() view returns(uint256)
func (_Api *ApiCallerSession) Getbalance0() (*big.Int, error) {
	return _Api.Contract.Getbalance0(&_Api.CallOpts)
}

// Getbalance1 is a free data retrieval call binding the contract method 0x385602f8.
//
// Solidity: function getbalance1() view returns(uint256)
func (_Api *ApiCaller) Getbalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getbalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Getbalance1 is a free data retrieval call binding the contract method 0x385602f8.
//
// Solidity: function getbalance1() view returns(uint256)
func (_Api *ApiSession) Getbalance1() (*big.Int, error) {
	return _Api.Contract.Getbalance1(&_Api.CallOpts)
}

// Getbalance1 is a free data retrieval call binding the contract method 0x385602f8.
//
// Solidity: function getbalance1() view returns(uint256)
func (_Api *ApiCallerSession) Getbalance1() (*big.Int, error) {
	return _Api.Contract.Getbalance1(&_Api.CallOpts)
}

// IndividualCap is a free data retrieval call binding the contract method 0x0276650b.
//
// Solidity: function individualCap() view returns(uint256)
func (_Api *ApiCaller) IndividualCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "individualCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndividualCap is a free data retrieval call binding the contract method 0x0276650b.
//
// Solidity: function individualCap() view returns(uint256)
func (_Api *ApiSession) IndividualCap() (*big.Int, error) {
	return _Api.Contract.IndividualCap(&_Api.CallOpts)
}

// IndividualCap is a free data retrieval call binding the contract method 0x0276650b.
//
// Solidity: function individualCap() view returns(uint256)
func (_Api *ApiCallerSession) IndividualCap() (*big.Int, error) {
	return _Api.Contract.IndividualCap(&_Api.CallOpts)
}

// IsUsed is a free data retrieval call binding the contract method 0x5ccaf589.
//
// Solidity: function isUsed() view returns(bool)
func (_Api *ApiCaller) IsUsed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isUsed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUsed is a free data retrieval call binding the contract method 0x5ccaf589.
//
// Solidity: function isUsed() view returns(bool)
func (_Api *ApiSession) IsUsed() (bool, error) {
	return _Api.Contract.IsUsed(&_Api.CallOpts)
}

// IsUsed is a free data retrieval call binding the contract method 0x5ccaf589.
//
// Solidity: function isUsed() view returns(bool)
func (_Api *ApiCallerSession) IsUsed() (bool, error) {
	return _Api.Contract.IsUsed(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCallerSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_Api *ApiCaller) Strategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "strategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_Api *ApiSession) Strategy() (common.Address, error) {
	return _Api.Contract.Strategy(&_Api.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_Api *ApiCallerSession) Strategy() (common.Address, error) {
	return _Api.Contract.Strategy(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCallerSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Api *ApiCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Api *ApiSession) Token0() (common.Address, error) {
	return _Api.Contract.Token0(&_Api.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Api *ApiCallerSession) Token0() (common.Address, error) {
	return _Api.Contract.Token0(&_Api.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Api *ApiCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Api *ApiSession) Token1() (common.Address, error) {
	return _Api.Contract.Token1(&_Api.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Api *ApiCallerSession) Token1() (common.Address, error) {
	return _Api.Contract.Token1(&_Api.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiCallerSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// VaultCap is a free data retrieval call binding the contract method 0x3c1bda09.
//
// Solidity: function vaultCap() view returns(uint256)
func (_Api *ApiCaller) VaultCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "vaultCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultCap is a free data retrieval call binding the contract method 0x3c1bda09.
//
// Solidity: function vaultCap() view returns(uint256)
func (_Api *ApiSession) VaultCap() (*big.Int, error) {
	return _Api.Contract.VaultCap(&_Api.CallOpts)
}

// VaultCap is a free data retrieval call binding the contract method 0x3c1bda09.
//
// Solidity: function vaultCap() view returns(uint256)
func (_Api *ApiCallerSession) VaultCap() (*big.Int, error) {
	return _Api.Contract.VaultCap(&_Api.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Api *ApiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Api *ApiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, spender, amount)
}

// CalcPositionShares is a paid mutator transaction binding the contract method 0x33e2f2ea.
//
// Solidity: function calcPositionShares(uint256 amountIn0, uint256 amountIn1) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiTransactor) CalcPositionShares(opts *bind.TransactOpts, amountIn0 *big.Int, amountIn1 *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "calcPositionShares", amountIn0, amountIn1)
}

// CalcPositionShares is a paid mutator transaction binding the contract method 0x33e2f2ea.
//
// Solidity: function calcPositionShares(uint256 amountIn0, uint256 amountIn1) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiSession) CalcPositionShares(amountIn0 *big.Int, amountIn1 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CalcPositionShares(&_Api.TransactOpts, amountIn0, amountIn1)
}

// CalcPositionShares is a paid mutator transaction binding the contract method 0x33e2f2ea.
//
// Solidity: function calcPositionShares(uint256 amountIn0, uint256 amountIn1) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiTransactorSession) CalcPositionShares(amountIn0 *big.Int, amountIn1 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CalcPositionShares(&_Api.TransactOpts, amountIn0, amountIn1)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Api *ApiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Api *ApiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.DecreaseAllowance(&_Api.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Api *ApiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.DecreaseAllowance(&_Api.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 amountIn0, uint256 amountIn1) returns()
func (_Api *ApiTransactor) Deposit(opts *bind.TransactOpts, amountIn0 *big.Int, amountIn1 *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "deposit", amountIn0, amountIn1)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 amountIn0, uint256 amountIn1) returns()
func (_Api *ApiSession) Deposit(amountIn0 *big.Int, amountIn1 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountIn0, amountIn1)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 amountIn0, uint256 amountIn1) returns()
func (_Api *ApiTransactorSession) Deposit(amountIn0 *big.Int, amountIn1 *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, amountIn0, amountIn1)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xd8d7f96f.
//
// Solidity: function emergencyBurn() returns()
func (_Api *ApiTransactor) EmergencyBurn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "emergencyBurn")
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xd8d7f96f.
//
// Solidity: function emergencyBurn() returns()
func (_Api *ApiSession) EmergencyBurn() (*types.Transaction, error) {
	return _Api.Contract.EmergencyBurn(&_Api.TransactOpts)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xd8d7f96f.
//
// Solidity: function emergencyBurn() returns()
func (_Api *ApiTransactorSession) EmergencyBurn() (*types.Transaction, error) {
	return _Api.Contract.EmergencyBurn(&_Api.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Api *ApiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Api *ApiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.IncreaseAllowance(&_Api.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Api *ApiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Api.Contract.IncreaseAllowance(&_Api.TransactOpts, spender, addedValue)
}

// MoveFunds is a paid mutator transaction binding the contract method 0xd0b6ba2a.
//
// Solidity: function moveFunds() returns()
func (_Api *ApiTransactor) MoveFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "moveFunds")
}

// MoveFunds is a paid mutator transaction binding the contract method 0xd0b6ba2a.
//
// Solidity: function moveFunds() returns()
func (_Api *ApiSession) MoveFunds() (*types.Transaction, error) {
	return _Api.Contract.MoveFunds(&_Api.TransactOpts)
}

// MoveFunds is a paid mutator transaction binding the contract method 0xd0b6ba2a.
//
// Solidity: function moveFunds() returns()
func (_Api *ApiTransactorSession) MoveFunds() (*types.Transaction, error) {
	return _Api.Contract.MoveFunds(&_Api.TransactOpts)
}

// SetIndividualCap is a paid mutator transaction binding the contract method 0xcd2b3026.
//
// Solidity: function setIndividualCap(uint256 newMax) returns()
func (_Api *ApiTransactor) SetIndividualCap(opts *bind.TransactOpts, newMax *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setIndividualCap", newMax)
}

// SetIndividualCap is a paid mutator transaction binding the contract method 0xcd2b3026.
//
// Solidity: function setIndividualCap(uint256 newMax) returns()
func (_Api *ApiSession) SetIndividualCap(newMax *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetIndividualCap(&_Api.TransactOpts, newMax)
}

// SetIndividualCap is a paid mutator transaction binding the contract method 0xcd2b3026.
//
// Solidity: function setIndividualCap(uint256 newMax) returns()
func (_Api *ApiTransactorSession) SetIndividualCap(newMax *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetIndividualCap(&_Api.TransactOpts, newMax)
}

// SetSrategy is a paid mutator transaction binding the contract method 0xd46915bc.
//
// Solidity: function setSrategy(address _strat) returns()
func (_Api *ApiTransactor) SetSrategy(opts *bind.TransactOpts, _strat common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setSrategy", _strat)
}

// SetSrategy is a paid mutator transaction binding the contract method 0xd46915bc.
//
// Solidity: function setSrategy(address _strat) returns()
func (_Api *ApiSession) SetSrategy(_strat common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetSrategy(&_Api.TransactOpts, _strat)
}

// SetSrategy is a paid mutator transaction binding the contract method 0xd46915bc.
//
// Solidity: function setSrategy(address _strat) returns()
func (_Api *ApiTransactorSession) SetSrategy(_strat common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetSrategy(&_Api.TransactOpts, _strat)
}

// SetVaultCap is a paid mutator transaction binding the contract method 0xf0f5907d.
//
// Solidity: function setVaultCap(uint256 newMax) returns()
func (_Api *ApiTransactor) SetVaultCap(opts *bind.TransactOpts, newMax *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setVaultCap", newMax)
}

// SetVaultCap is a paid mutator transaction binding the contract method 0xf0f5907d.
//
// Solidity: function setVaultCap(uint256 newMax) returns()
func (_Api *ApiSession) SetVaultCap(newMax *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetVaultCap(&_Api.TransactOpts, newMax)
}

// SetVaultCap is a paid mutator transaction binding the contract method 0xf0f5907d.
//
// Solidity: function setVaultCap(uint256 newMax) returns()
func (_Api *ApiTransactorSession) SetVaultCap(newMax *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetVaultCap(&_Api.TransactOpts, newMax)
}

// Sweep is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address _token, uint256 amount) returns()
func (_Api *ApiTransactor) Sweep(opts *bind.TransactOpts, _token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sweep", _token, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address _token, uint256 amount) returns()
func (_Api *ApiSession) Sweep(_token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Sweep(&_Api.TransactOpts, _token, amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x6ea056a9.
//
// Solidity: function sweep(address _token, uint256 amount) returns()
func (_Api *ApiTransactorSession) Sweep(_token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Sweep(&_Api.TransactOpts, _token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Api *ApiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Api *ApiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Api *ApiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc6ab5d90.
//
// Solidity: function withdraw(uint8 percent) returns(uint256 amount0, uint256 amount1)
func (_Api *ApiTransactor) Withdraw(opts *bind.TransactOpts, percent uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "withdraw", percent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc6ab5d90.
//
// Solidity: function withdraw(uint8 percent) returns(uint256 amount0, uint256 amount1)
func (_Api *ApiSession) Withdraw(percent uint8) (*types.Transaction, error) {
	return _Api.Contract.Withdraw(&_Api.TransactOpts, percent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc6ab5d90.
//
// Solidity: function withdraw(uint8 percent) returns(uint256 amount0, uint256 amount1)
func (_Api *ApiTransactorSession) Withdraw(percent uint8) (*types.Transaction, error) {
	return _Api.Contract.Withdraw(&_Api.TransactOpts, percent)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Api *ApiTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Api.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Api *ApiSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Api.Contract.Fallback(&_Api.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Api *ApiTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Api.Contract.Fallback(&_Api.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactorSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

// ApiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Api contract.
type ApiApprovalIterator struct {
	Event *ApiApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiApproval represents a Approval event raised by the Api contract.
type ApiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Api *ApiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ApiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ApiApprovalIterator{contract: _Api.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Api *ApiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ApiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiApproval)
				if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Api *ApiFilterer) ParseApproval(log types.Log) (*ApiApproval, error) {
	event := new(ApiApproval)
	if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Api contract.
type ApiDepositIterator struct {
	Event *ApiDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiDeposit represents a Deposit event raised by the Api contract.
type ApiDeposit struct {
	Sender  common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*ApiDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiDepositIterator{contract: _Api.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ApiDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiDeposit)
				if err := _Api.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x36af321ec8d3c75236829c5317affd40ddb308863a1236d2d277a4025cccee1e.
//
// Solidity: event Deposit(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) ParseDeposit(log types.Log) (*ApiDeposit, error) {
	event := new(ApiDeposit)
	if err := _Api.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Api contract.
type ApiTransferIterator struct {
	Event *ApiTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiTransfer represents a Transfer event raised by the Api contract.
type ApiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Api *ApiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ApiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ApiTransferIterator{contract: _Api.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Api *ApiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ApiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiTransfer)
				if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Api *ApiFilterer) ParseTransfer(log types.Log) (*ApiTransfer, error) {
	event := new(ApiTransfer)
	if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Api contract.
type ApiWithdrawIterator struct {
	Event *ApiWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWithdraw represents a Withdraw event raised by the Api contract.
type ApiWithdraw struct {
	Sender  common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address) (*ApiWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiWithdrawIterator{contract: _Api.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ApiWithdraw, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWithdraw)
				if err := _Api.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed sender, uint256 shares, uint256 amount0, uint256 amount1)
func (_Api *ApiFilterer) ParseWithdraw(log types.Log) (*ApiWithdraw, error) {
	event := new(ApiWithdraw)
	if err := _Api.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

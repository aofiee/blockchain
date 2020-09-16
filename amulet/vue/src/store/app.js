module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "buying", fields: ["AmuletHash", "BuyingPrice", "SellingBy", "BuyingBy", ] },
		{ type: "selling", fields: ["AmuletHash", "SellingPrice", "SellingBy", ] },
		{ type: "amulet", fields: ["Name", "Identity", "Address", "Telephon", "AmuletType", "AmuletName", "From", "Model", "Surface", "Year", "Province", "Description", "Remark", ] },
  ],
};

package storage

import (
	"errors"

	"gorm.io/gorm"
)


type AppOps struct {
	App *App
}

func (app *AppOps) Info() *App {
	return app.App
}


func (app *AppOps) Update(kwargs map[string]interface{}) {
	app.App.Name = setDefault(kwargs, "Name", app.App.Name)
	app.App.Description = setDefault(kwargs, "Description", app.App.Description)
	app.App.Icon = setDefault(kwargs, "Icon", app.App.Icon)
	app.App.Mode = setDefault(kwargs, "Mode", app.App.Mode)
	app.App.Path = setDefault(kwargs, "Path", app.App.Path)
	app.App.Order = setDefault(kwargs, "Order", app.App.Order)
	app.App.RoleID = setDefault(kwargs, "RoleID", app.App.RoleID)
	app.App.CategoryID = setDefault(kwargs, "CategoryID", app.App.CategoryID)
	DBClient.Save(app.App)
}


func (app *AppOps) Delete() {
	DBClient.Delete(app.App)
}


func CreateApp(kwargs map[string]interface{}) (*App, error) {
	app := App{
		Name: setDefault(kwargs, "Name", ""),
		Description: setDefault(kwargs, "Description", ""),
		Icon: setDefault(kwargs, "Icon", ""),
		Mode: setDefault(kwargs, "Mode", InnerMode),
		Path: setDefault(kwargs, "Path", ""),
		Order: setDefault(kwargs, "Order", 1),
		RoleID:  setDefault(kwargs, "RoleID", RegularUserID),
		CategoryID: setDefault(kwargs, "CategoryID", DefaultID),
	}
	result := DBClient.Create(&app)
	if result.Error != nil {
		return nil, result.Error
	}
	return &app, nil
}


func GetApps(userRole int) ([]*App, error) {
	var apps []*App
	result := DBClient.Where("role_id <= ?", userRole).Find(&apps).Order("category_id")
	if result.Error != nil {
		return nil, result.Error
	}
	return apps, nil
}


func GetAppByID(ID int) (*App, error) {
	var app = App{ID: ID}
	result := DBClient.First(&app)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &app, nil
}



func setDefault[T any](kwargs map[string]interface{}, key string, defaultValue T) T {
	value, ok := kwargs[key]
	if ok {
		if typedValue, ok := value.(T); ok {
			return typedValue
		}
	}
	return value.(T)
}


// NewAppOps creates a new AppOps object.
//
// Parameters:
// - name: the name of the app.
// - isCreate: a boolean indicating whether to create a new app.
// - kwargs: a map of key-value pairs containing additional arguments.
//
// Returns:
// - *AppOps: a pointer to the created AppOps object.
// - error: an error if any occurred during the creation process.
func NewAppOps(id int) (*AppOps, error) {
	var app = App{ID: id}
	result := DBClient.First(&app)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &AppOps{App: &app}, nil
}


func RegisterCategories() AppCategory {
	var systemCategory = AppCategory{
		ID: SystemID,
		Icon:  "",
		Name:  SystemName,
		Order: SystemOrder,
	}
	var defaultCategory = AppCategory{
		ID: DefaultID,
		Icon:  "",
		Name:  DefaultName,
		Order: DefaultOrder,
	}
	result := DBClient.Create(&systemCategory)
	if result.Error != nil {
		panic(result.Error)
	}
	result = DBClient.Create(&defaultCategory)
	if result.Error != nil {
		panic(result.Error)
	}
	return systemCategory
}


// RegisterApps is a function that performs a pre-dump of the apps.
//
// It does the following:
// - Retrieves the app category from the database.
// - If the category is not found, it calls the PreDumpCategories function.
// - Creates two apps: "应用中心" and "权限中心" with specific properties.
// - Inserts the apps into the database.
//
// No parameters are required.
// No return values.
func RegisterApps() {
	var category = AppCategory{Name: SystemName}
	result := DBClient.First(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		category = RegisterCategories()
	}
	var appCenter = App{
		ID: AppCenterID,
		Order:	     0,
		Name:        "应用中心",
		Description: "应用管理中心",
		Icon:        "",
		Mode:        InnerMode,
		Path:        "/system/app-center",
		RoleID:        SystemAdminID,
		CategoryID:  category.ID,
	}
	var roleCentor = App{
		ID: RoleCenterID,
		Order: 		 1,
		Name:        "权限中心",
		Description: "权限管理中心",
		Icon:        "",
		Mode:        InnerMode,
		Path:        "/system/role-center",
		RoleID:        SystemAdminID,
		CategoryID:  category.ID,

	}
	result = DBClient.First(&appCenter)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		DBClient.Create(&appCenter)
	}
	result = DBClient.First(&roleCentor)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		DBClient.Create(&roleCentor)
	}
}



#include "main.hpp"

namespace lab32 {
    MainView::MainView() {
        set_title("Beer paint");
        set_default_size(200, 200);

        set_child(canvas_);
        canvas_.set_content_height(100);
        canvas_.set_content_width(100);
    }
    MainView::~MainView() {}
}


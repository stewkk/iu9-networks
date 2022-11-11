#ifndef MAIN_VIEW_H_
#define MAIN_VIEW_H_

#include <gtkmm.h>

#include "canvas.hpp"

namespace lab32 {
    class MainView : public Gtk::Window {
        public:
            MainView();
            virtual ~MainView();
        protected:
            Canvas canvas_;
    };
}

#endif // MAIN_VIEW_H_
